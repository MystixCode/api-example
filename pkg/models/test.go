package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"

	"time"
)

type Test struct {
	ID          string    `json:"id" gorm:"primarykey"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

type TestBeforeSave struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

func (c *Test) GetAll(db *gorm.DB) (*[]Test, error) {
	var err error
	var tests []Test

	err = db.Find(&tests).Error
	if err != nil {
		return nil, err
	}

	return &tests, nil
}

func (c *Test) Save(db *gorm.DB, newTest *TestBeforeSave) error {
	var err error

	id := uuid.New().String()
	err = db.Where(&Test{ID: id}).First(c).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			c.ID = id
		} else {
			return err
		}
	}

	c.Name = newTest.Name
	c.Description = newTest.Description
	c.CreatedAt = time.Now()
	c.UpdatedAt = time.Now()

	err = db.Create(c).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *Test) GetByID(db *gorm.DB, id string) error {
	var err error

	err = db.Where(&Test{ID: id}).First(c).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *Test) Update(db *gorm.DB, id string, newTest *TestBeforeSave) error {
	var err error

	err = c.GetByID(db, id)

	if newTest.Name != "" {
		c.Name = newTest.Name
	}

	if newTest.Description != "" {
		c.Description = newTest.Description
	}

	c.UpdatedAt = time.Now()

	err = db.Where(&Test{ID: c.ID}).Updates(c).Error
	if err != nil {
		return err
	}

	return nil
}

func (c *Test) Delete(db *gorm.DB, id string) error {
	var err error

	err = db.Where(&Test{ID: id}).Delete(c).Error
	if err != nil {
		return err
	}

	return nil
}
