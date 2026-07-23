package product

import "gorm.io/gorm"

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{db: db}
}

func (r *Repository) Create(product *Product) error {
	return r.db.Create(product).Error
}

func (r *Repository) GetByID(id uint) (*Product, error) {
	var product Product
	if err := r.db.First(&product, id).Error; err != nil {
		return nil, err
	}
	return &product, nil
}

func (r *Repository) GetAll() ([]Product, error) {
	var products []Product
	if err := r.db.Find(&products).Error; err != nil {
		return nil, err
	}
	return products, nil
}

func (r *Repository) Update(product *Product) (*Product, error) {
	if err := r.db.Save(product).Error; err != nil {
		return nil, err
	}
	return product, nil
}

func (r *Repository) Delete(id uint) error {
	if err := r.db.Delete(&Product{}, id).Error; err != nil {
		return err
	}
	return nil
}
