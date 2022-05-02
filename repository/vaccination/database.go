package vaccination

import (
	"context"
	"errors"
	"fmt"

	"github.com/daffashafwan/vaxin-service/business/vaccinations"
	"gorm.io/gorm"
)

type VaccinationRepo struct {
	DB *gorm.DB
}

func CreateVaccinationRepo(conn *gorm.DB) vaccinations.Repository {
	return &VaccinationRepo{
		DB: conn,
	}
}

func (rep *VaccinationRepo) Create(ctx context.Context, eventCreate *vaccinations.Domain) (vaccinations.Domain, error) {
	event := Vaccination{
		EventId:   eventCreate.EventId,
		UserId:    eventCreate.UserId,
		SpecialId: eventCreate.SpecialId,
		Queue:     eventCreate.Queue,
	}
	err := rep.DB.Create(&event)
	if err.Error != nil {
		return vaccinations.Domain{}, err.Error
	}
	errs := rep.DB.Table("vaccinations").Preload("Event.Vaccine").Preload("Event.Facility").Preload("User").Find(&event, "id=?", event.Id)
	if errs.Error != nil {
		return vaccinations.Domain{}, errs.Error
	}
	fmt.Println(event)
	return event.ToDomain(), nil
}

func (repo *VaccinationRepo) Update(ctx context.Context, eventUpdate vaccinations.Domain) (vaccinations.Domain, error) {
	data := FromDomain(eventUpdate)
	err := repo.DB.Table("vaccinations").First(&data)
	if err.Error != nil {
		return vaccinations.Domain{}, err.Error
	}
	data.EventId = eventUpdate.EventId
	data.UserId = eventUpdate.UserId
	data.SpecialId = eventUpdate.SpecialId
	data.Queue = eventUpdate.Queue
	if repo.DB.Save(&data).Error != nil {
		return vaccinations.Domain{}, errors.New("bad requests")
	}

	return data.ToDomain(), nil
}

func (repo *VaccinationRepo) GetAll(ctx context.Context) ([]vaccinations.Domain, error) {
	var data []Vaccination
	err := repo.DB.Table("vaccinations").Preload("Event").Preload("User").Find(&data)
	if err.Error != nil {
		return []vaccinations.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (repo *VaccinationRepo) GetById(ctx context.Context, id int) (vaccinations.Domain, error) {
	var data Vaccination
	err := repo.DB.Table("vaccinations").Preload("Event").Preload("User").Find(&data, "id=?", id)
	if err.Error != nil {
		return vaccinations.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}

func (repo *VaccinationRepo) GetBySpecialId(ctx context.Context, id string) (vaccinations.Domain, error) {
	var data Vaccination
	err := repo.DB.Table("vaccinations").Find(&data, "special_id=?", id)
	if err.Error != nil {
		return vaccinations.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}

func (repo *VaccinationRepo) GetByEventId(ctx context.Context, id int) ([]vaccinations.Domain, error) {
	var data []Vaccination
	err := repo.DB.Table("vaccinations").Preload("Event").Preload("User").Find(&data, "event_id=?", id)
	if err.Error != nil {
		return []vaccinations.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (repo *VaccinationRepo) GetByUserId(ctx context.Context, id int) ([]vaccinations.Domain, error) {
	var data []Vaccination
	err := repo.DB.Table("vaccinations").Preload("Event").Preload("User").Find(&data, "user_id=?", id)
	if err.Error != nil {
		return []vaccinations.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (repo *VaccinationRepo) Delete(ctx context.Context, id int) error {
	event := Vaccination{}
	err := repo.DB.Table("vaccinations").Where("id = ?", id).First(&event).Delete(&event)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}