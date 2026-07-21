package product

import (
	"4-order-api/pkg"
	"encoding/json"
	"net/http"

	"github.com/go-playground/validator/v10"
	"gorm.io/gorm"
)

type HandlerDeps struct {
	DB *gorm.DB
}
type Handler struct {
	deps HandlerDeps
}

func NewHandler(deps HandlerDeps) *Handler {
	return &Handler{deps: deps}
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var payload CreateProductPayload
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		pkg.JsonResponse(w, http.StatusBadRequest, pkg.SendResponse{Message: "Invalid request payload"})
		return
	}
	if err := validator.New().Struct(payload); err != nil {
		pkg.JsonResponse(w, http.StatusBadRequest, pkg.SendResponse{Message: "Invalid request payload"})
		return
	}
	p := &Product{
		Name:        payload.Name,
		Description: payload.Description,
		Images:      payload.Images,
	}
	if err := h.deps.DB.Create(p).Error; err != nil {
		pkg.JsonResponse(w, http.StatusInternalServerError, pkg.SendResponse{Message: "Failed to create product"})
		return
	}
	pkg.JsonResponse(w, http.StatusCreated, pkg.SendResponse{Message: "Product created successfully"})

}
