package infrastructure

import (
	"github.com/labstack/echo/v4"

	"shukai-api/interfaces/action"
)

type Routing struct {
	SqlHandler *SqlHandler
	Echo       *echo.Echo
	Port       string
}

func NewRouting(sqlHandler *SqlHandler, port string) *Routing {
	r := &Routing{
		SqlHandler: sqlHandler,
		Echo:       echo.New(),
		Port:       port,
	}
	r.setRouting()
	return r
}

func (r *Routing) setRouting() {
	decideAction := func(c echo.Context) error { return action.NewApprovalDecideAction(r.SqlHandler).Decide(c) }
	userAddAction := func(c echo.Context) error { return action.NewUserAddAction(r.SqlHandler).Add(c) }
	userEditAction := func(c echo.Context) error { return action.NewUserEditAction(r.SqlHandler).Put(c) }
	userGetAction := func(c echo.Context) error { return action.NewUserGetAction(r.SqlHandler).Get(c) }
	recruitmentAddAction := func(c echo.Context) error { return action.NewRecruitmentAddAction(r.SqlHandler).Add(c) }
	recruitmentGetListAction := func(c echo.Context) error { return action.NewRecruitmentGetListAction(r.SqlHandler).GetList(c) }
	r.Echo.DELETE("/approval/:approvalwait_id/:should_approval", decideAction)
	r.Echo.POST("/user/add", userAddAction)
	r.Echo.PUT("/user/:user_id", userEditAction)
	r.Echo.GET("/user/:user_id", userGetAction)
	// r.Echo.PATCH("/notification/:user_id/:notification_id")
	r.Echo.GET("/recruitment/:user_id", recruitmentGetListAction)
	// r.Echo.GET("/recruitment/:recruitment_id")
	r.Echo.POST("/recruitment/add", recruitmentAddAction)
	// r.Echo.POST("/recruitment/request/:recruitment_id")
	// r.Echo.DELETE("/recruitment/cancel/:recruitment_id/:user_id")
	// r.Echo.POST("/user/add")
	// r.Echo.POST("/login")
}

func (r *Routing) Run() {
	r.setRouting()
	r.Echo.Logger.Fatal(r.Echo.Start(":" + r.Port))
}
