package controller

import (
	"go-api/model"
	"go-api/usecase"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type productController struct {
	productUsecase usecase.ProductUsecase
}

func NewProductController(useCase usecase.ProductUsecase) productController {
	return productController{
		productUsecase: useCase,
	}
}

func (p *productController) GetProducts(ctx *gin.Context) {

	products, err := p.productUsecase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
	}

	ctx.JSON(http.StatusOK, products)
}

func (p *productController) CreateProduct(ctx *gin.Context) {
	var product model.Product

	err := ctx.ShouldBind(&product)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "JSON inválido"})
		return
	}

	idProductCreated, err := p.productUsecase.CreateProduct(product)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{"idProduct": idProductCreated})
}

func (p *productController) GetProductById(ctx *gin.Context) {
	id := ctx.Param("product-id")

	if id == "" {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "Id do produto não pode ser vazio"})
		return
	}

	idProduct, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "Id do produto deve ser um número"})
		return
	}

	product, err := p.productUsecase.GetProductById(idProduct)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if product == nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"erro": "Produto não encontrado"})
		return
	}

	ctx.JSON(http.StatusOK, product)
}
