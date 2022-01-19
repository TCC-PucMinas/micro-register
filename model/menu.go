package model

type Menus struct {
	Id         int64      `json:"id"`
	Name       string     `json:"name"`
	IconMenu   string     `json:"icon_menu"`
	UrlMenu    string     `json:"url_menu"`
	Permission Permission `json:"permissio"`
}
