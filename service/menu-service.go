package service

import (
	"log"
	"rumah-makan/dto"
	"rumah-makan/model"
	"rumah-makan/repository"

	"github.com/mashingan/smapping"
)

type MenuService interface {
	Insert(m dto.MenuCreateDTO) model.Menu
	Update(m dto.MenuUpdateDTO) model.Menu
	Delete(m model.Menu)
	All() []model.Menu
	FindMenuByID(menuID uint64) model.Menu
	// IsAllowedToEdit(customerID string, menuID uint64) bool
	InsertImage(ID uint64, fileLocation string) model.Menu
}

type menuService struct {
	menuRepository repository.MenuRepository
}

func NewMenuService(menuRepo repository.MenuRepository) MenuService {
	return &menuService{
		menuRepository: menuRepo,
	}
}

func (service *menuService) Insert(m dto.MenuCreateDTO) model.Menu {
	menu := model.Menu{}
	err := smapping.FillStruct(&menu, smapping.MapFields(&m))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.menuRepository.InsertMenu(menu)
	return res
}

func (service *menuService) Update(m dto.MenuUpdateDTO) model.Menu {
	menu := model.Menu{}
	err := smapping.FillStruct(&menu, smapping.MapFields(&m))
	if err != nil {
		log.Fatalf("Failed map %v: ", err)
	}
	res := service.menuRepository.UpdateMenu(menu)
	return res
}

func (service *menuService) Delete(m model.Menu) {
	service.menuRepository.DeleteMenu(m)
}

func (service *menuService) All() []model.Menu {
	return service.menuRepository.AllMenu()
}

func (service *menuService) FindMenuByID(menuID uint64) model.Menu {
	return service.menuRepository.FindMenuByID(menuID)
}

func (service *menuService) InsertImage(ID uint64, fileLocation string) model.Menu {
	menu := service.menuRepository.FindByID(ID)
	menu.Image = fileLocation
	res := service.menuRepository.InsertMenuImage(menu)
	return res
}


// func (service *menuService) IsAllowedToEdit(customerID string, menuID uint64) bool {
// 	m := service.menuRepository.FindMenuByID(menuID)
// 	id := fmt.Sprintf("%v", m.CustomerID)
// 	return customerID == id
// }