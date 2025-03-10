package handler

import (
	"net/http"
	"strconv"
	usecase "test-kp-golang/src/domain/product/use-case"
	"test-kp-golang/src/utils"

	"github.com/gin-gonic/gin"
)

type ProductHandler struct {
	productUsecase *usecase.ProductUsecase
}

func NewProductHandler(r *gin.Engine, productUsecase *usecase.ProductUsecase) {
	handler := &ProductHandler{productUsecase: productUsecase}

	r.GET("/products", handler.GetProducts)
	r.GET("/products/:id", handler.GetProductByID)
}

func (h *ProductHandler) GetProducts(c *gin.Context) {
	products, err := h.productUsecase.GetProducts()
	if err != nil {
		utils.FormatErrorResponse(c, "error", http.StatusBadRequest, err.Error(), nil)
		return
	}
	utils.FormatErrorResponse(c, "success", http.StatusOK, "success", products)
}

func (h *ProductHandler) GetProductByID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		utils.FormatErrorResponse(c, "error", http.StatusBadRequest, err.Error(), nil)
		return
	}
	product, err := h.productUsecase.GetProductByID(id)
	if err != nil {
		utils.FormatErrorResponse(c, "error", http.StatusBadRequest, err.Error(), nil)
		return
	}

	utils.FormatErrorResponse(c, "success", http.StatusOK, "success", product)
}
