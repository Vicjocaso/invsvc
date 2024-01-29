package handlers

import (
	"fmt"
	"invsvc/repos"
	"invsvc/types"
	"net/http"

	"github.com/labstack/echo/v4"
)

type ProductHanlder struct {
	prodcutsRepo repos.ProductRepoInterface
}

func NewProductHandler(prodcutsRepo *repos.ProductRepo) *ProductHanlder {
	return &ProductHanlder{
		prodcutsRepo: prodcutsRepo,
	}
}

func (h *ProductHanlder) GetAll(c echo.Context) error {

	products, err := h.prodcutsRepo.GetProduct()
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %v", err))
	}

	return c.JSON(http.StatusOK, products)
}

func (h *ProductHanlder) GetBy(c echo.Context) error {
	productId := c.Param("id")
	product, err := h.prodcutsRepo.GetById(productId)
	if err != nil {
		c.String(http.StatusInternalServerError, fmt.Sprintf("Error: %v", err))
	}

	return c.JSON(http.StatusOK, product)
}

func (h *ProductHanlder) Create(c echo.Context) error {
	newProduct := new(types.Product)
	if err := c.Bind(newProduct); err != nil {
		return err
	}
	return c.JSON(http.StatusCreated, newProduct)
}
