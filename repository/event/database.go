package event

import (
	"context"
	"errors"
	"time"
	"github.com/daffashafwan/vaxin-service/business/events"
	"gorm.io/gorm"
)

type EventRepo struct {
	DB *gorm.DB
}

func CreateEventRepo(conn *gorm.DB) events.Repository {
	return &EventRepo{
		DB: conn,
	}
}

func (rep *EventRepo) Create(ctx context.Context, eventCreate *events.Domain) (events.Domain, error) {
	layout := "2006-01-02T15:04:05.000Z"
	start, _ := time.Parse(layout, eventCreate.StartDate)
	end, _ := time.Parse(layout, eventCreate.EndDate)
	event := Event{
		FacilityId: eventCreate.FacilityId,
		VaccineId:  eventCreate.VaccineId,
		Name:       eventCreate.Name,
		Quota:      eventCreate.Quota,
		StartDate:  start,
		EndDate:    end,
	}
	err := rep.DB.Create(&event)
	if err.Error != nil {
		return events.Domain{}, err.Error
	}
	return event.ToDomain(), nil
}

func (repo *EventRepo) Update(ctx context.Context, eventUpdate events.Domain) (events.Domain, error) {
	data := FromDomain(eventUpdate)
	err := repo.DB.Table("events").First(&data)
	if err.Error != nil {
		return events.Domain{}, err.Error
	}
	layout := "2006-01-02T15:04:05.000Z"
	start, _ := time.Parse(layout, eventUpdate.StartDate)
	end, _ := time.Parse(layout, eventUpdate.EndDate)
	data.FacilityId = eventUpdate.FacilityId
	data.VaccineId = eventUpdate.VaccineId
	data.Name = eventUpdate.Name
	data.Quota = eventUpdate.Quota
	data.StartDate = start
	data.EndDate = end
	if repo.DB.Save(&data).Error != nil {
		return events.Domain{}, errors.New("bad requests")
	}

	return data.ToDomain(), nil
}

func (repo *EventRepo) UpdateQuota(ctx context.Context, eventUpdate events.Domain) (events.Domain, error) {
	data := FromDomain(eventUpdate)
	err := repo.DB.Table("events").First(&data)
	if err.Error != nil {
		return events.Domain{}, err.Error
	}
	data.Quota = eventUpdate.Quota
	if repo.DB.Save(&data).Error != nil {
		return events.Domain{}, errors.New("bad requests")
	}

	return data.ToDomain(), nil
}

func (repo *EventRepo) UpdateQueue(ctx context.Context, eventUpdate events.Domain) (events.Domain, error) {
	data := FromDomain(eventUpdate)
	err := repo.DB.Table("events").First(&data)
	if err.Error != nil {
		return events.Domain{}, err.Error
	}
	data.Queue = eventUpdate.Queue
	if repo.DB.Save(&data).Error != nil {
		return events.Domain{}, errors.New("bad requests")
	}
	return data.ToDomain(), nil
}

func (repo *EventRepo) GetAll(ctx context.Context) ([]events.Domain, error) {
	var data []Event
	err := repo.DB.Table("events").Preload("Vaccine").Preload("Facility").Find(&data)
	if err.Error != nil {
		return []events.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (repo *EventRepo) GetById(ctx context.Context, id int) (events.Domain, error) {
	var data Event
	err := repo.DB.Table("events").Preload("Vaccine").Preload("Facility").Find(&data, "id=?", id)
	if err.Error != nil {
		return events.Domain{}, err.Error
	}
	return data.ToDomain(), nil
}

func (repo *EventRepo) GetByVaccineId(ctx context.Context, id int) ([]events.Domain, error) {
	var data []Event
	err := repo.DB.Table("events").Preload("Vaccine").Preload("Facility").Find(&data, "vaccine_id=?", id)
	if err.Error != nil {
		return []events.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (repo *EventRepo) GetByFacilityId(ctx context.Context, id int) ([]events.Domain, error) {
	var data []Event
	err := repo.DB.Table("events").Preload("Vaccine").Preload("Facility").Find(&data, "facility_id=?", id)
	if err.Error != nil {
		return []events.Domain{}, err.Error
	}
	return ToListDomain(data), nil
}

func (repo *EventRepo) Delete(ctx context.Context, id int) error {
	event := Event{}
	err := repo.DB.Table("events").Where("id = ?", id).First(&event).Delete(&event)
	if err.Error != nil {
		return err.Error
	}
	if err.RowsAffected == 0 {
		return errors.New("id not found")
	}
	return nil
}
