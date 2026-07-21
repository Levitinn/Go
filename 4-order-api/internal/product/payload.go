package product

import "github.com/lib/pq"

type CreateProductPayload struct {
	Name        string         `json:"name" validate:"required"`
	Description string         `json:"description" validate:"required"`
	Images      pq.StringArray `json:"images" validate:"required"`
}
