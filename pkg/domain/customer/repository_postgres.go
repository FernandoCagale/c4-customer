package customer

import (
	"fmt"
	"github.com/FernandoCagale/c4-customer/internal/errors"
	"github.com/FernandoCagale/c4-customer/pkg/entity"
	"github.com/jinzhu/gorm"
)

type GormRepository struct {
	db *gorm.DB
}

func NewGormRepository(db *gorm.DB) *GormRepository {
	return &GormRepository{db}
}

func (repo *GormRepository) FindAll() (customers []*entity.Customer, err error) {
	if err := repo.db.Find(&customers).Error; err != nil {
		return nil, errors.ErrInternalServer
	}

	return customers, nil
}

func (repo *GormRepository) FindById(ID string) (customer *entity.Customer, err error) {
	var customerr entity.Customer

	if err := repo.db.First(&customerr, ID).Error; err != nil {
		fmt.Println(err.Error())
		switch err {
		case gorm.ErrRecordNotFound:
			return nil, errors.ErrNotFound
		default:
			return nil, errors.ErrInternalServer
		}
	}

	return &customerr, nil
}

func (repo *GormRepository) DeleteById(ID string) (err error) {
	var customer entity.Customer

	if err := repo.db.First(&customer, ID).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return errors.ErrNotFound
		default:
			return errors.ErrInternalServer
		}
	}

	if err := repo.db.Model(&customer).Delete(customer).Error; err != nil {
		return errors.ErrInternalServer
	}

	return  nil
}

func (repo *GormRepository) Create(e *entity.Customer) (err error) {
	if err := repo.db.Create(&e).Error; err != nil {
		switch err {
		case gorm.ErrRecordNotFound:
			return errors.ErrNotFound
		default:
			return errors.ErrInternalServer
		}
	}
	return nil
}
