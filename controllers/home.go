package controllers
import (
	"goweb1/model"
    "net/http"
	"html/template"
	"github.com/julienschmidt/httprouter"
)

type (  
    HomeController struct{}
)

func (hc HomeController) Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {  
	products, _ := model.GetAllProduct()
	categories, _ := model.GetAllCategory()

	data := map[string]interface{}{
		"Products": products,
		"Categories":  categories,
	}
	// layout file must be the first parameter in ParseFiles!
	templates, err := template.ParseFiles(
		"views/layout/master.html",
		"views/layout/header.html",
		"views/home/index.html",
		"views/layout/footer.html",	
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := templates.Execute(w, data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}