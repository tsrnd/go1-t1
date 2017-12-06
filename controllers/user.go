package controllers
import (
	"goweb1/model"
	"strconv"
    "net/http"
	"html/template"
	"github.com/julienschmidt/httprouter"
	"github.com/gorilla/csrf"
)

type (  
    UserController struct{}
)

func (uc UserController) GetUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {  
	
	id := ps.ByName("id")
	ids, err := strconv.ParseInt(id, 10, 64)
	user, _ := model.GetUserByID(ids)
	if (err != nil )|| (len(user) == 0) {
		http.Redirect(w, r, URL_HOME, http.StatusMovedPermanently)
	} else {
		session, _ := store.Get(r, "session-id")
		if session.Values["username"] == nil ||  session.Values["id"] != user[0].ID {
			http.Redirect(w, r, URL_NOTFOUND, http.StatusMovedPermanently)
		}
	}
	session, _ := store.Get(r, "session-id")
	username := session.Values["username"]
	cats, _ := model.GetAllCategory()
	
	udata := map[string]interface{}{
	"cats":         cats,
	"name":       	username,
	"user":			user,
	csrf.TemplateTag: csrf.TemplateField(r),
	}
	
	// layout file must be the first parameter in ParseFiles!
	templates, err := template.ParseFiles(
		"views/layout/master.html",	
		"views/layout/header.html",
		"views/user/view.html",
		"views/layout/footer.html",
	)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := templates.Execute(w, udata); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (uc UserController) UpdateUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	id := ps.ByName("id")
	ids, _ := strconv.ParseInt(id, 10, 64)
	user, _ := model.GetUserByID(ids)

	r.ParseForm()
	old_pasword := r.FormValue("old-password")
	new_pasword := r.FormValue("new-pasword")
	fullname := r.FormValue("fullname")
	address := r.FormValue("address")
	password,_ := HashPassword(new_pasword)
	
	check := CheckPasswordHash(old_pasword,user[0].Password)
	if check == false || new_pasword  == ""  {
		http.Redirect(w, r, "/user/"+id, http.StatusMovedPermanently)
	} else {
		model.UpdateUser(ids, fullname, address, password)
		http.Redirect(w, r, URL_HOME, http.StatusMovedPermanently)
	}
}