package repository

import (
	"github.com/Dparty/dao"
	"gorm.io/gorm"
)

type Repository[T any] struct {
	db *gorm.DB
}

func NewRepository[T any]() *Repository[T] {
	return &Repository[T]{
		db: dao.GetDBInstance(),
	}
}

func (r Repository[T]) Find(dest T, conds ...any) {
	r.db.Find(dest, conds...)
}

func (r Repository[T]) GetById(id uint) *T {
	var dest T
	ctx := r.db.Find(&dest, id)
	if ctx.RowsAffected == 0 {
		return nil
	}
	return &dest
}

func (r Repository[T]) Create(dest T) {
	r.db.Create(dest)
}

func (r Repository[T]) Update(dest T) {
	r.db.Save(dest)
}

func (r Repository[T]) Delete(dest T) {
	r.db.Delete(dest)
}
