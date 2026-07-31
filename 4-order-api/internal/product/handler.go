package product

import (
	"4-order-api/pkg"
	"net/http"
)

type HandlerDeps struct {
	Repository *Repository
}
type Handler struct {
	deps HandlerDeps
}

func NewHandler(deps HandlerDeps) *Handler {
	return &Handler{deps: deps}
}

func (h *Handler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var payload CreateProductPayload
	if err := pkg.DecodeJSON(r, &payload); err != nil {
		pkg.JsonResponse(w, http.StatusBadRequest, pkg.SendResponse{Message: "Invalid request payload"})
		return
	}
	p := &Product{
		Name:        payload.Name,
		Description: payload.Description,
		Images:      payload.Images,
	}
	if err := h.deps.Repository.Create(p); err != nil {
		pkg.JsonResponse(w, http.StatusInternalServerError, pkg.SendResponse{Message: "Failed to create product"})
		return
	}
	pkg.JsonResponse(w, http.StatusCreated, p)
}

func (h *Handler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id, err := pkg.ParseID(r)
	if err != nil {
		pkg.JsonResponse(w, http.StatusBadRequest, pkg.SendResponse{Message: "Invalid product ID"})
		return
	}
	var payload UpdateProductPayload
	if err := pkg.DecodeJSON(r, &payload); err != nil {
		pkg.JsonResponse(w, http.StatusBadRequest, pkg.SendResponse{Message: "Invalid request payload"})
		return
	}
	existing, err := h.deps.Repository.GetByID(uint(id))
	if err != nil {
		pkg.JsonResponse(w, http.StatusNotFound, pkg.SendResponse{Message: "Product not found"})
		return
	}
	existing.Name = payload.Name
	existing.Description = payload.Description
	existing.Images = payload.Images
	updated, err := h.deps.Repository.Update(existing)
	if err != nil {
		pkg.JsonResponse(w, http.StatusInternalServerError, pkg.SendResponse{Message: "Failed to update product"})
		return
	}
	pkg.JsonResponse(w, http.StatusOK, updated)
}

func (h *Handler) GetAllProducts(w http.ResponseWriter, r *http.Request) {
	products, err := h.deps.Repository.GetAll()
	if err != nil {
		pkg.JsonResponse(w, http.StatusInternalServerError, pkg.SendResponse{Message: "Failed to get products"})
		return
	}
	pkg.JsonResponse(w, http.StatusOK, products)
}

func (h *Handler) GetProductByID(w http.ResponseWriter, r *http.Request) {
	id, err := pkg.ParseID(r)
	if err != nil {
		pkg.JsonResponse(w, http.StatusBadRequest, pkg.SendResponse{Message: "Invalid product ID"})
		return
	}
	product, err := h.deps.Repository.GetByID(uint(id))
	if err != nil {
		pkg.JsonResponse(w, http.StatusNotFound, pkg.SendResponse{Message: "Product not found"})
		return
	}
	pkg.JsonResponse(w, http.StatusOK, product)
}

func (h *Handler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id, err := pkg.ParseID(r)
	if err != nil {
		pkg.JsonResponse(w, http.StatusBadRequest, pkg.SendResponse{Message: "Invalid product ID"})
		return
	}
	existing, err := h.deps.Repository.GetByID(id)
	if err != nil {
		pkg.JsonResponse(w, http.StatusNotFound, pkg.SendResponse{Message: "Product not found"})
		return
	}
	if err := h.deps.Repository.Delete(existing.ID); err != nil {
		pkg.JsonResponse(w, http.StatusInternalServerError, pkg.SendResponse{Message: "Failed to delete product"})
		return
	}
	pkg.JsonResponse(w, http.StatusOK, pkg.SendResponse{Message: "Product deleted successfully"})
}
