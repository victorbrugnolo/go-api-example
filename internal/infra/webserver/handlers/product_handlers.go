package handlers

import (
	"encoding/json"
	"github.com/go-chi/chi"
	"github.com/victorbrugnolo/go-api-example/internal/dto"
	"github.com/victorbrugnolo/go-api-example/internal/entity"
	"github.com/victorbrugnolo/go-api-example/internal/infra/database"
	entityPkg "github.com/victorbrugnolo/go-api-example/pkg/entity"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	ProductDB database.ProductInterface
}

func NewProductHandler(db database.ProductInterface) *ProductHandler {
	return &ProductHandler{
		ProductDB: db,
	}
}

// CreateProduct godoc
// @Summary      Create product
// @Description  Create products
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        request     body      dto.CreateProductInput  true  "product request"
// @Success      201
// @Failure      500         {object}  Error
// @Router       /products [post]
// @Security ApiKeyAuth
func (h *ProductHandler) CreateProduct(w http.ResponseWriter, r *http.Request) {
	var product dto.CreateProductInput

	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	p, err := entity.NewProduct(product.Name, product.Price)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = h.ProductDB.Create(p)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (h *ProductHandler) GetProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductDB.FindByID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(product)
}

// GetProducts godoc
// @Summary      List products
// @Description  get all products
// @Tags         products
// @Accept       json
// @Produce      json
// @Param        page      query     string  false  "page number"
// @Param        limit     query     string  false  "limit"
// @Success      200       {array}   entity.Product
// @Failure      404       {object}  Error
// @Failure      500       {object}  Error
// @Router       /products [get]
// @Security ApiKeyAuth
func (h *ProductHandler) GetProducts(w http.ResponseWriter, r *http.Request) {
	page := r.URL.Query().Get("page")
	limit := r.URL.Query().Get("limit")
	sort := r.URL.Query().Get("sort")

	pageInt, err := strconv.Atoi(page)

	if err != nil {
		pageInt = 0
	}

	limitInt, err := strconv.Atoi(limit)

	if err != nil {
		limitInt = 0
	}

	products, err := h.ProductDB.FindAll(pageInt, limitInt, sort)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(products)
}

func (h *ProductHandler) UpdateProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var product entity.Product
	err := json.NewDecoder(r.Body).Decode(&product)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product.ID, err = entityPkg.ParseID(id)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = h.ProductDB.FindByID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Update(&product)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (h *ProductHandler) DeleteProduct(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	product, err := h.ProductDB.FindByID(id)

	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = h.ProductDB.Delete(product)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
