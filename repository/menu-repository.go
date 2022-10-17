package repository

import (
	"rumah-makan/model"

	"gorm.io/gorm"
)

type MenuRepository interface {
	InsertMenu(m model.Menu) model.Menu
	UpdateMenu(m model.Menu) model.Menu
	DeleteMenu(m model.Menu)
	AllMenu() []model.Menu
	FindMenuByID(menuID uint64) model.Menu
}

type menuConnection struct {
	connection *gorm.DB
}

func NewMenuRepository(dbConn *gorm.DB) MenuRepository {
	return &menuConnection{
		connection: dbConn,
	}
}

func (db *menuConnection) InsertMenu(m model.Menu) model.Menu {
	db.connection.Save(&m)
	return m
}

func (db *menuConnection) UpdateMenu(m model.Menu) model.Menu {
	db.connection.Save(&m)
	return m
}

func (db *menuConnection) DeleteMenu(m model.Menu) {
	db.connection.Delete(&m)
}

func (db *menuConnection) AllMenu() []model.Menu {
	var menus []model.Menu
	db.connection.Find(&menus)
	return menus
}

func (db *menuConnection) FindMenuByID(menuID uint64) model.Menu {
	var menu model.Menu
	db.connection.Preload("Customer").Find(&menu, menuID)
	return menu
}