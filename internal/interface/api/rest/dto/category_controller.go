package rest

import (
	"net/http"

	app_interfaces "github.com/Cthulhu-tech/golang_blog/internal/application/interface"
	domain_interfaces "github.com/Cthulhu-tech/golang_blog/internal/domain/interface"
	controller_response_mapper "github.com/Cthulhu-tech/golang_blog/internal/interface/api/rest/dto/mapper"
	"github.com/google/uuid"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

type CategoryController struct {
	service app_interfaces.CategoryService
}

func NewCategoryController(e *echo.Echo, service domain_interfaces.CategoryRepository) *CategoryController {
	controller := &CategoryController{
		service: service,
	}

	e.POST("/api/v1/products", controller.service.CreateCategory)
	e.GET("/api/v1/products", controller.service.GetAllCategories)
	e.GET("/api/v1/products/:id", controller.service.GetCategoryByID)
	e.Use(middleware.Recover())

	return controller
}

// func (cc *CategoryController) CreateCategoryController(c echo.Context) error {
// 	var createCategoryRequest request.CreateProductRequest

// 	if err := c.Bind(&createProductRequest); err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{
// 			"error": "Failed to parse request body",
// 		})
// 	}

// 	productCommand, err := createCategoryRequest.ToCreateProductCommand()
// 	if err != nil {
// 		return c.JSON(http.StatusBadRequest, map[string]string{
// 			"error": "Invalid product Id format",
// 		})
// 	}

// 	result, err := cc.service.CreateCategory(productCommand)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{
// 			"error": "Failed to create product",
// 		})
// 	}

// 	response := db_mapper.ToProductResponse(result.Result)

// 	return c.JSON(http.StatusCreated, response)
// }

// func (cc *CategoryController) GetAllCategoriesController(c echo.Context) error {
// 	products, err := cc.service.GetAllCategories(1, 25)
// 	if err != nil {
// 		return c.JSON(http.StatusInternalServerError, map[string]string{
// 			"error": "Failed to fetch products",
// 		})
// 	}

// 	response := controller_response_mapper.ToCategoryResponse(products.Result)

// 	return c.JSON(http.StatusOK, response)
// }

func (cc *CategoryController) GetCategoryByIdController(c echo.Context) error {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid product Id format",
		})
	}

	product, err := cc.service.GetCategoryByID(id.String())
	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to fetch product",
		})
	}

	if product == nil {
		return c.JSON(http.StatusNotFound, map[string]string{
			"error": "Product not found",
		})
	}

	response := controller_response_mapper.ToCategoryResponse(product.Result)

	return c.JSON(http.StatusOK, response)
}
