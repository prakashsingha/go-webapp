package controller

import (
	"html/template"
	"net/http"
	"regexp"
	"strconv"

	"github.com/prakashsingha/go-webapp/model"
	"github.com/prakashsingha/go-webapp/viewmodel"
)

type shop struct {
	shopTemplate     *template.Template
	categoryTemplate *template.Template
}

func (sh shop) registerRoutes() {
	http.HandleFunc("/shop", sh.shopHandler)
	http.HandleFunc("/shop/", sh.shopHandler)
}

func (sh shop) shopHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "text/html")

	categoryPattern, _ := regexp.Compile(`/shop/(\d+)`)
	matches := categoryPattern.FindStringSubmatch(r.URL.Path)
	if len(matches) > 0 {
		categoryID, _ := strconv.Atoi(matches[1])
		sh.handleCategory(w, r, categoryID)
	} else {
		categories := model.GetCategories()
		vm := viewmodel.NewShop(categories)
		sh.shopTemplate.Execute(w, vm)
	}
}

func (sh shop) handleCategory(w http.ResponseWriter, r *http.Request, categoryID int) {
	products := model.GetProductsForCategory(categoryID)
	vm := viewmodel.NewShopDetail(products)
	sh.categoryTemplate.Execute(w, vm)
}
