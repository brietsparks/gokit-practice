package abilities

import (
	"context"
	"github.com/satori/go.uuid"
	"github.com/jinzhu/gorm"
	//"github.com/davecgh/go-spew/spew"
	"errors"
)

type Ability struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key"`
	OwnerId uuid.UUID `gorm:"type:uuid;primary_key"`
	Caption string
}

type Service interface {
	CreateAbility(ctx context.Context, a Ability) (Ability, error)
	UpdateAbility(ctx context.Context, a Ability) (Ability, error)
	DeleteAbility(ctx context.Context, a Ability) error

	//GetAbility(ctx context.Context, a Ability) (Ability, error)
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &service{
		db: db,
	}
}

func (s *service) CreateAbility(ctx context.Context, a Ability) (Ability, error) {
	db := s.db.Create(&a)

	return a, db.Error
}

func (s *service) UpdateAbility(ctx context.Context, a Ability) (Ability, error) {
	db := s.db.Save(&a)

	return a, db.Error
}

func (s *service) DeleteAbility(ctx context.Context, a Ability) (error) {
	db := s.db.Delete(&a)

	if db.Error != nil {
		return db.Error
	}

	if db.RowsAffected == 0 {
		return errors.New("trying to delete non existent Ability")
	}

	return nil
}

//func (s *service) GetAbility(ctx context.Context, a Ability) (Ability, error) {
//	db := s.db.First(&a)
//
//	return a, db.Error
//}