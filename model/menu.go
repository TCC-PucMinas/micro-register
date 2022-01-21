package model

import (
	"strconv"

	"github.com/TCC-PucMinas/micro-register/db"
)

type Menus struct {
	Id         int64      `json:"id"`
	Name       string     `json:"name"`
	IconMenu   string     `json:"icon_menu"`
	UrlMenu    string     `json:"url_menu"`
	Permission Permission `json:"permission"`
}

func GetAllMenus(permissionId int64) ([]Menus, error) {

	sql := db.ConnectDatabase()

	menusArray := []Menus{}

	m := Menus{}

	query := `select id, name, icon_menu, url_menu from menus where id_permission = ?`

	requestConfig, err := sql.Query(query, permissionId)

	if err != nil {
		return menusArray, err
	}

	for requestConfig.Next() {
		var id, name, icon_menu, url_menu string
		_ = requestConfig.Scan(&id, &name, &icon_menu, &url_menu)
		i64, _ := strconv.ParseInt(id, 10, 64)
		m.Id = i64
		m.Name = name
		m.IconMenu = icon_menu
		m.UrlMenu = url_menu

		menusArray = append(menusArray, m)
	}

	return menusArray, nil
}
