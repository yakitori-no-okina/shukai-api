package infrastructure

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

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
	// r.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
	// 	AllowOrigins: []string{"https://syukai.netlify.app/"},
	// 	AllowMethods: []string{http.MethodGet, http.MethodPatch, http.MethodPut, http.MethodPost, http.MethodDelete},
	// }))
	r.Echo.Use(middleware.CORS())

	decideAction := func(c echo.Context) error { return action.NewApprovalDecideAction(r.SqlHandler).Decide(c) }
	notificationGetListAction := func(c echo.Context) error { return action.NewNotificationGetListAction(r.SqlHandler).GetList(c) }
	notificationReadAction := func(c echo.Context) error { return action.NewNotificationReadAction(r.SqlHandler).Read(c) }
	recruitmentAddAction := func(c echo.Context) error { return action.NewRecruitmentAddAction(r.SqlHandler).Add(c) }
	recruitmentGetListAction := func(c echo.Context) error { return action.NewRecruitmentGetListAction(r.SqlHandler).GetList(c) }
	recruitmentGetAction := func(c echo.Context) error { return action.NewRecruitmentGetAction(r.SqlHandler).Get(c) }
	recruitmentRequestAction := func(c echo.Context) error { return action.NewRecruitmentRequestAction(r.SqlHandler).Request(c) }
	recruitmentCancelAction := func(c echo.Context) error { return action.NewRecruitmentCancelAction(r.SqlHandler).Cancel(c) }
	userAddAction := func(c echo.Context) error { return action.NewUserAddAction(r.SqlHandler).Add(c) }
	userEditAction := func(c echo.Context) error { return action.NewUserEditAction(r.SqlHandler).Put(c) }
	userGetAction := func(c echo.Context) error { return action.NewUserGetAction(r.SqlHandler).Get(c) }
	loginAction := func(c echo.Context) error { return action.NewLoginAction(r.SqlHandler).Login(c) }

	r.Echo.DELETE("/approval/:approvalwait_id/:should_approval", decideAction)
	r.Echo.GET("/notification/:user_id", notificationGetListAction)
	r.Echo.PATCH("/notification/read/:notification_id", notificationReadAction)
	r.Echo.GET("/recruitment/:user_id/list", recruitmentGetListAction)
	r.Echo.GET("/recruitment/:recruitment_id", recruitmentGetAction)
	r.Echo.POST("/recruitment/add", recruitmentAddAction)
	r.Echo.POST("/recruitment/request/:recruitment_id", recruitmentRequestAction)
	r.Echo.DELETE("/recruitment/cancel/:recruitment_id/:user_id", recruitmentCancelAction)
	r.Echo.POST("/user/add", userAddAction)
	r.Echo.PUT("/user/:user_id", userEditAction)
	r.Echo.GET("/user/:user_id", userGetAction)
	r.Echo.POST("/login", loginAction)
}

func (r *Routing) Run() {
	r.setRouting()
	r.Echo.Logger.Fatal(r.Echo.Start(":" + r.Port))
}
