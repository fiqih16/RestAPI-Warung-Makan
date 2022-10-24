package service

import (
	"rumah-makan/dto"
	"rumah-makan/model"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockRepository struct {
	mock.Mock
}

// AllMenu implements repository.MenuRepository
func (*MockRepository) AllMenu() []model.Menu {
	panic("unimplemented")
}

// DeleteMenu implements repository.MenuRepository
func (*MockRepository) DeleteMenu(m model.Menu) {
	panic("unimplemented")
}

// FindByID implements repository.MenuRepository
func (*MockRepository) FindByID(ID uint64) model.Menu {
	panic("unimplemented")
}

// FindMenuByID implements repository.MenuRepository
func (*MockRepository) FindMenuByID(menuID uint64) model.Menu {
	panic("unimplemented")
}

// InsertMenu implements repository.MenuRepository
func (mock *MockRepository) InsertMenu(m model.Menu) model.Menu {
	args := mock.Called(m)
	return args.Get(0).(model.Menu)
}

// InsertMenuImage implements repository.MenuRepository
func (*MockRepository) InsertMenuImage(m model.Menu) model.Menu {
	panic("unimplemented")
}

// UpdateMenu implements repository.MenuRepository
func (*MockRepository) UpdateMenu(m model.Menu) model.Menu {
	panic("unimplemented")
}

func TestInsert(t *testing.T) {
	mockRepo := new(MockRepository)
	menu := model.Menu{
		NamaMenu: "Nasi Goreng",
		Harga:    10000,
		Status:   "Tersedia",
	}

	mockRepo.On("InsertMenu", menu).Return(menu)

	menuService := NewMenuService(mockRepo)
	result := menuService.Insert(dto.MenuCreateDTO{
		NamaMenu: "Nasi Goreng",
		Harga:    10000,
		Status:  "Tersedia",
	})

	assert.Equal(t, "Nasi Goreng", result.NamaMenu)
	assert.Equal(t, 10000, result.Harga)
	assert.Equal(t, "Tersedia", result.Status)
	mockRepo.AssertExpectations(t)

	
}
