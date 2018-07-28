package abilities

import (
	"context"
	"github.com/satori/go.uuid"
	"github.com/jinzhu/gorm"
	"gokit-practice/util"
)

type Ability struct {
	Id      uuid.UUID `gorm:"type:uuid;primary_key"`
	OwnerId uuid.UUID `gorm:"type:uuid;primary_key"`
	Caption string
}

type Service interface {
	GetAbilitiesByOwnerId(ctx context.Context, ownerId string) ([]Ability, error)
	CreateAbility(ctx context.Context, a Ability) (Ability, error)
	UpdateAbility(ctx context.Context, a Ability) (Ability, error)
	DeleteAbility(ctx context.Context, a Ability) (Ability, error)
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) *service {
	return &service{
		db: db,
	}
}

const (
	AbilityDNE     = "ABILITY_DNE"
)

func (s *service) GetAbilitiesByOwnerId(ctx context.Context, ownerId string) ([]Ability, error) {
	var a []Ability
	s.db.Where("owner_id = ?", ownerId).Find(&a)

	if s.db.Error != nil {
		return a, util.NewUnexpectedError(s.db.Error)
	}

	return a, nil
}

func (s *service) CreateAbility(ctx context.Context, a Ability) (Ability, error) {
	db := s.db.Create(&a)

	if db.Error != nil {
		return a, util.NewUnexpectedError(db.Error)
	}

	return a, nil
}

func (s *service) UpdateAbility(ctx context.Context, a Ability) (Ability, error) {
	db := s.db.Save(&a)

	if db.Error != nil {
		return a, util.NewUnexpectedError(db.Error)
	}

	return a, nil
}

func (s *service) DeleteAbility(ctx context.Context, a Ability) (Ability, error) {
	db := s.db.Delete(&a)

	if db.Error != nil {
		return a, util.NewUnexpectedError(db.Error)
	}

	if db.RowsAffected == 0 {
		return a, util.NewError(AbilityDNE, "trying to delete non-existent Ability")
	}

	return a, nil
}
