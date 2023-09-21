package utility

import (
	"CSIE_ADMS_Back-End/internal/model/entity"
	"github.com/gogf/gf/v2/frame/g"
)

func IsAdmin(id int) bool {
	user := entity.Users{}
	m := g.Model("users").Safe()
	err := m.Scan(&user, "id", id)
	if err != nil {
		return false
	}
	return user.Role == "admin"
}
