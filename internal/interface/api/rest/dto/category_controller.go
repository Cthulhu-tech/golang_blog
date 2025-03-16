package rest

import (
	"net/http"

	app_interfaces "github.com/Cthulhu-tech/golang_blog/internal/application/interface"
	controller_response_mapper "github.com/Cthulhu-tech/golang_blog/internal/interface/api/rest/dto/mapper"
	"github.com/Cthulhu-tech/golang_blog/internal/interface/api/rest/dto/request"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type CategoryController struct {
	service app_interfaces.CategoryService
}

func NewCategoryController(e *echo.Echo, service app_interfaces.CategoryService) *CategoryController {
	controller := &CategoryController{
		service: service,
	}

	e.POST("/api/v1/category", controller.CreateCategory)
	e.GET("/api/v1/category", controller.GetAllCategories)
	e.GET("/api/v1/category/:id", controller.GetCategoryByID)
	e.Use(middleware.Recover())

	return controller
}

func (cc *CategoryController) CreateCategory(c echo.Context) error {
	var createCategoryRequest request.CreateCategoryRequest

	if err := c.Bind(&createCategoryRequest); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to parse request body",
		})
	}

	categoryCommand, err := createCategoryRequest.ToCreateCategoryCommand()
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid category Id format",
		})
	}

	result, err := cc.service.CreateCategory(categoryCommand)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create category",
		})
	}

	response := controller_response_mapper.ToCategoryResponse(result.Result)

	return c.JSON(http.StatusCreated, response)
}

func (cc *CategoryController) GetAllCategories(c echo.Context) error {
	categories, err := cc.service.GetAllCategories(1, 25)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch Categories",
		})
	}

	if categories == nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Categories not found",
		})
	}

	response := controller_response_mapper.ToCategoryListResponse(categories.Result)

	return c.JSON(http.StatusOK, response)
}

func (cc *CategoryController) GetCategoryByID(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid category Id format",
		})
	}

	product, err := cc.service.GetCategoryByID(id.String())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch category",
		})
	}

	if product == nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Category not found",
		})
	}

	response := controller_response_mapper.ToCategoryResponse(product.Result)

	return c.JSON(http.StatusOK, response)
}
