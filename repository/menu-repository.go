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
	FindByID(ID uint64) model.Menu
	InsertMenuImage(m model.Menu) model.Menu
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
	db.connection.Exec("INSERT INTO menus (nama_menu, harga, status, image) VALUES (?, ?, ?, ?)", m.NamaMenu, m.Harga, m.Status, m.Image)
	return m
}

func (db *menuConnection) UpdateMenu(m model.Menu) model.Menu {
	db.connection.Exec("UPDATE menus SET nama_menu = ?, harga = ?, status = ? WHERE id = ?", m.NamaMenu, m.Harga, m.Status, m.ID)
	db.connection.Find(&m)
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
	db.connection.Find(&menu, menuID)
	return menu
}

func (db *menuConnection) FindByID(ID uint64) model.Menu {
	var menu model.Menu
	db.connection.Where("id = ?", ID).Find(&menu)
	return menu
}

func (db *menuConnection) InsertMenuImage(m model.Menu) model.Menu {
	db.connection.Exec("UPDATE menus SET image = ? WHERE id = ?", m.Image, m.ID)
	db.connection.Find(&m)
	return m
}