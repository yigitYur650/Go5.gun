// Package dal provides data access layer for todo items.
package dal

type Todo struct {
	ID           int // sütun numaralarını giriyor bu go da defult bişey
	TamamlandıMI string
	tamam        bool `gorm:"default:false"`
}
