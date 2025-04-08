package repositories

import (
	"gorm.io/gorm"
)

type GenericRepositoryImpl[Tentity any] struct {
	db *gorm.DB
}

func NewGenericRepository[Tentity any](db *gorm.DB) *GenericRepositoryImpl[Tentity] {
	return &GenericRepositoryImpl[Tentity]{db: db}
}
func (r *GenericRepositoryImpl[Tentity]) GetAll(filters *Tentity, preload []string) ([]Tentity, error) {
	var entities []Tentity
	query := r.db
	for _, p := range preload {
		query = query.Preload(p)
	}

	if filters != nil {
		query = query.Where(filters)
	}

	if result := query.Find(&entities); result.Error != nil {
		return nil, result.Error
	}

	return entities, nil
}
func (r *GenericRepositoryImpl[Tentity]) Get(filters *Tentity, preload []string) (*Tentity, error) {
	var entity *Tentity
	query := r.db
	for _, p := range preload {
		query = query.Preload(p)
	}

	if filters != nil {
		query = query.Where(filters)
	}

	if result := query.First(&entity); result.Error != nil {
		return nil, result.Error
	}
	return entity, nil
}
func (r *GenericRepositoryImpl[Tentity]) Add(entity *Tentity) error {
	if result := r.db.Create(&entity); result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *GenericRepositoryImpl[Tentity]) Update(entity *Tentity) error {
	if result := r.db.Save(&entity); result.Error != nil {
		return result.Error
	}
	return nil
}
func (r *GenericRepositoryImpl[Tentity]) Delete(entity *Tentity) error {
	if result := r.db.Delete(&entity); result.Error != nil {
		return result.Error
	}
	return nil
}
