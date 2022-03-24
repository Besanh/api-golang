package v1

import (
	"api-golang/service"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	userService service.IUser
}

func APIUser(r *gin.Engine, u service.IUser) {
	user := &UserHandler{
		userService: u,
	}
	// Group cac route chung middleware
	group := r.Group("v1/user")
	{
		group.GET("", user.GetUsers) // GET All
	}
}

func (h *UserHandler) GetUsers(g *gin.Context) {
	limit := g.Query("limit")
	offset := g.Query("offset")
	limmitInt, err := strconv.Atoi(limit)
	if err != nil {
		limmitInt = 50
	}
	offsetInt, err := strconv.Atoi(offset)
	if err != nil {
		offsetInt = 0
	}
	code, result := h.userService.GetUser(g, limmitInt, offsetInt)
	g.JSON(code, result)
}
