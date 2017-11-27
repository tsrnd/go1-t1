package controllers
import (
    "net/http"
	"html/template"
	"github.com/julienschmidt/httprouter"
	// "first_webapp/database"
	// "first_webapp/model"
	//"fmt"
)

type (  
    HomeController struct{}
)

type Content struct {
    Title string
}

func (hc HomeController) Home(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {  

	context := Content{"Home Page"}
	
	// layout file must be the first parameter in ParseFiles!
	templates, err := template.ParseFiles(
		"views/layout/master.html",
		"views/layout/header.html",
		"views/layout/slider.html",
		"views/home/index.html",
		"views/layout/footer.html",
		
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := templates.Execute(w, context); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}