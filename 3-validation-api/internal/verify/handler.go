package verify

import (
	"3-validation-api/config"
	"3-validation-api/internal"
	responses "3-validation-api/packages"
	"encoding/json"
	"io"
	"net/http"
	"net/smtp"

	"github.com/google/uuid"
	"github.com/jordan-wright/email"
)

type HandlerDeps struct {
	Config *config.Config
}

type Handler struct {
	config *config.Config
	tokens map[string]string
}

func NewHandler(deps HandlerDeps) *Handler {
	return &Handler{config: deps.Config, tokens: make(map[string]string)}
}

func (handler *Handler) Send(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer r.Body.Close()

	var req internal.SendPayload
	err = json.Unmarshal(body, &req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if req.Email == "" {
		http.Error(w, "email is required", http.StatusBadRequest)
		return
	}
	secret := uuid.New().String()
	handler.tokens[secret] = req.Email
	e := email.NewEmail()
	e.From = handler.config.Email
	e.To = []string{req.Email}
	e.Subject = "Verification Code"
	e.Text = []byte("http://localhost:8081/verify/" + secret)
	e.HTML = []byte("<p>Your verification code is: <strong>" + secret + "</strong></p>")
	err = e.Send(
		handler.config.Address,
		smtp.PlainAuth("", handler.config.Email, handler.config.Password, "smtp.gmail.com"),
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	responses.JsonResponse(w, http.StatusOK, "ok")

}

func (handler *Handler) Verify(w http.ResponseWriter, r *http.Request) {
	hash := r.PathValue("hash")
	email, ok := handler.tokens[hash]
	if !ok {
		responses.JsonResponse(w, http.StatusBadRequest, responses.SendResponse{Message: "invalid hash"})
		return
	}

	responses.JsonResponse(w, http.StatusOK, responses.SendResponse{Message: email})
}
