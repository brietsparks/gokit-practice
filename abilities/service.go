package abilities

import (
	"github.com/satori/go.uuid"
	"github.com/jinzhu/gorm"
	"gokit-practice/util"
)

type Ability struct {
	Id      uuid.UUID `gorm:"type:uuid;primary_key"`
	OwnerId uuid.UUID `gorm:"type:uuid;primary_key"`
	Caption string
}

type serviceWriteMethod func(Ability) (Ability, error)

type Service interface {
	GetAbilitiesByOwnerId(ownerId string) ([]Ability, error)
	CreateAbility(a Ability) (Ability, error)
	UpdateAbility(a Ability) (Ability, error)
	DeleteAbility(a Ability) (Ability, error)
}

type service struct {
	db *gorm.DB
}

func NewService(db *gorm.DB) Service {
	return &service{
		db: db,
	}
}

const (
	AbilityDNE     = "ABILITY_DNE"
)

func (s *service) GetAbilitiesByOwnerId(ownerId string) ([]Ability, error) {
	var a []Ability
	s.db.Where("owner_id = ?", ownerId).Find(&a)

	if s.db.Error != nil {
		return a, util.NewUnexpectedError(s.db.Error)
	}

	return a, nil
}

func (s *service) CreateAbility(a Ability) (Ability, error) {
	db := s.db.Create(&a)

	if db.Error != nil {
		return a, util.NewUnexpectedError(db.Error)
	}

	return a, nil
}

func (s *service) UpdateAbility(a Ability) (Ability, error) {
	db := s.db.Save(&a)

	if db.Error != nil {
		return a, util.NewUnexpectedError(db.Error)
	}

	return a, nil
}

func (s *service) DeleteAbility(a Ability) (Ability, error) {
	db := s.db.Delete(&a)

	if db.Error != nil {
		return a, util.NewUnexpectedError(db.Error)
	}

	if db.RowsAffected == 0 {
		return a, util.NewError(AbilityDNE, "trying to delete non-existent Ability")
	}

	return a, nil
}
