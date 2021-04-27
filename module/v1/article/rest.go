package article

import (
	"kumparan/model"
	"kumparan/module/v1/article/usecase"
	"kumparan/utl/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (m *Module) HandleRest(group *echo.Group) {
	group.GET("/list", m.articleList)
	group.POST("/create", m.ArticleNew)
}

// @Summary Get Article List
// @Description get all list article
// @Tags Article Management
// @Accept */*
// @Produce json
// @Success 200 {interface} model.Response{}
// @Router /api/v1/article/list [get]
func (m *Module) articleList(c echo.Context) error {
	var (
		requestId = c.Get("request_id").(string)
	)

	// usecase get user list
	resp, err := usecase.ArticleList(m.Config)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: err.Error(),
			Error:   err,
		})
	}

	return response.Success(c, model.Response{
		LogId:   requestId,
		Status:  http.StatusOK,
		Message: nil,
		Data:    resp,
	})
}

// @Summary Create Article
// @Description create article
// @Tags Article Management
// @Accept */*
// @Produce json
// @Success 200 {interface} model.Response{}
// @Router /api/v1/article/create [post]
func (m *Module) ArticleNew(c echo.Context) error {

	var (
		requestId = c.Get("request_id").(string)
		atr       = model.Article{}
	)

	err := c.Bind(&atr)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: nil,
			Error:   err,
		})
	}

	resp, err := usecase.ArticleNew(m.Config, &atr)
	if err != nil {
		return response.Error(c, model.Response{
			LogId:   requestId,
			Status:  http.StatusBadRequest,
			Message: "error",
			Error:   err.Error(),
		})
	}

	return response.Success(c, model.Response{
		LogId:   requestId,
		Status:  http.StatusOK,
		Message: "user has been registered",
		Data:    resp,
	})
}
