package article

import (
	"net/http"
	"strconv"

	"fr-circle-api/infrastructure"
	"fr-circle-api/shared/handler"

	"github.com/go-chi/chi"
	"github.com/sirupsen/logrus"
	validator "gopkg.in/go-playground/validator.v9"
)

// HTTPHandler struct.
type HTTPHandler struct {
	handler.BaseHTTPHandler
	aUsecase UsecaseInterface
}

// Get get all from article.
func (h *HTTPHandler) Get(w http.ResponseWriter, r *http.Request) {

	article, err := h.aUsecase.Get("")
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("h.aUsecase.Get errors")

		h.StatusServerError(w)
		return
	}

	// translation
	T := h.GetTranslaterFunc(r)

	common := CommonResponse{Result: 0, Message: T("success")}
	jsonData := GetResponse{CommonResponse: common, ResponseArticle: article}
	h.ResponseJSON(w, jsonData)
}

// GetID search id from article.
func (h *HTTPHandler) GetID(w http.ResponseWriter, r *http.Request) {
	// get url params
	id := chi.URLParam(r, "id")

	// validate get data.
	v := validator.New()
	err := v.Struct(GetIDRequest{ID: id})

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			h.Logger.Error("invalid validation error.")
			h.StatusServerError(w)
		}
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			// error example: https://github.com/go-playground/validator/blob/v9/_examples/struct-level/main.go
			errors = append(errors, err.StructField()+" is error.")
		}
		common := CommonResponse{Result: 1, Message: "", Errors: errors}
		h.StatusBadRequest(w, common)
		return
	}

	// change string to int
	iid, _ := strconv.Atoi(id)

	article, err := h.aUsecase.GetID(iid)
	if err != nil {
		T := h.GetTranslaterFunc(r)
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Fatal(T("error", map[string]interface{}{
			"Function": "h.aUsecase.GetID",
		}))
		h.StatusServerError(w)
		return
	}

	// translation
	T := h.GetTranslaterFunc(r)
	common := CommonResponse{Result: 0, Message: T("success")}
	jsonData := GetTitleResponse{CommonResponse: common, ResponseArticle: article}
	h.ResponseJSON(w, jsonData)
}

// PostAdd method.
func (h *HTTPHandler) PostAdd(w http.ResponseWriter, r *http.Request) {
	var err error
	// mapping post to struct.
	requestArticlePostAdd := PostAddRequest{}
	h.Parse(r, &requestArticlePostAdd)

	// log output.
	h.Logger.WithFields(logrus.Fields{
		"r.PostForm":      r.PostForm,
		"r.MultipartForm": r.MultipartForm,
	}).Debug("post data.")

	// log output.
	h.Logger.WithFields(logrus.Fields{
		"Title":   r.PostFormValue("title"),
		"Content": r.PostFormValue("content"),
	}).Debug("post data.")

	h.Logger.WithFields(logrus.Fields{
		"struct": requestArticlePostAdd,
	}).Debug("struct data.")

	// validate post data.
	v := validator.New()
	err = v.Struct(requestArticlePostAdd)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			h.Logger.Error("invalid validation error.")
			h.StatusServerError(w)
		}
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			// error example: https://github.com/go-playground/validator/blob/v9/_examples/struct-level/main.go
			errors = append(errors, err.StructField()+" is error.")
		}
		common := CommonResponse{Result: 1, Message: "", Errors: errors}
		h.StatusBadRequest(w, common)
		return
	}
	rowAffected, err := h.aUsecase.Add(requestArticlePostAdd.Title, requestArticlePostAdd.Content)

	if err != nil {
		h.StatusServerError(w)
		return
	}
	T := h.GetTranslaterFunc(r)
	common := CommonResponse{Result: 0, Message: T("add", map[string]interface{}{"Title": "article", "RowAffected": rowAffected})}
	h.ResponseJSON(w, common)
}

// DeleteID delete from article.
func (h *HTTPHandler) DeleteID(w http.ResponseWriter, r *http.Request) {
	var err error
	// get url params
	id := chi.URLParam(r, "id")

	// validate get data.
	v := validator.New()
	err = v.Struct(DeleteIDRequest{ID: id})

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			h.Logger.Error("invalid validation error.")
			h.StatusServerError(w)
		}
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			// error example: https://github.com/go-playground/validator/blob/v9/_examples/struct-level/main.go
			errors = append(errors, err.StructField()+" is error.")
		}
		common := CommonResponse{Result: 1, Message: "", Errors: errors}
		h.StatusBadRequest(w, common)
		return
	}

	// change string to int
	iid, _ := strconv.Atoi(id)

	rowAffected, err := h.aUsecase.Delete(iid)
	if err != nil {
		h.StatusServerError(w)
		return
	}

	T := h.GetTranslaterFunc(r)
	common := CommonResponse{Result: 0, Message: T("delete", map[string]interface{}{"Title": "article", "RowAffected": rowAffected})}
	h.ResponseJSON(w, common)
}

// GetCount get count from redis.
func (h *HTTPHandler) GetCount(w http.ResponseWriter, r *http.Request) {
	count, err := h.aUsecase.GetCount()
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("UsecaseInterface.GetCount() error.")
		h.StatusServerError(w)
		return
	}
	common := CommonResponse{Result: 0, Message: ""}
	jsonData := GetCountResponse{CommonResponse: common, Count: count}
	h.ResponseJSON(w, jsonData)
}

// PostCount get count from redis.
func (h *HTTPHandler) PostCount(w http.ResponseWriter, r *http.Request) {
	count, err := h.aUsecase.AddCount()
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("UsecaseInterface.GetCount() error.")
		h.StatusServerError(w)
		return
	}
	common := CommonResponse{Result: 0, Message: ""}
	jsonData := PostCountResponse{CommonResponse: common, Count: count}
	h.ResponseJSON(w, jsonData)
}

// PostVisenzeDiscoversearch requests visenze/discoversearch API
func (h *HTTPHandler) PostVisenzeDiscoversearch(w http.ResponseWriter, r *http.Request) {
	// mapping post to struct.
	request := PostVisenzeDiscoverSearchRequest{}
	h.ParseMultipart(r, &request)
	h.Logger.Debug(request)

	// file.
	f, fh, _ := r.FormFile("file")
	request.File = f
	// file close.
	if f != nil {
		defer f.Close()
	}

	// validate get data.
	v := validator.New()
	err := v.Struct(request)

	if err != nil {
		if _, ok := err.(*validator.InvalidValidationError); ok {
			h.Logger.Error("invalid validation error.")
			h.StatusServerError(w)
		}
		var errors []string
		for _, err := range err.(validator.ValidationErrors) {
			// error example: https://github.com/go-playground/validator/blob/v9/_examples/struct-level/main.go
			errors = append(errors, err.StructField()+" is error.")
		}
		common := CommonResponse{Result: 1, Message: "", Errors: errors}
		h.StatusBadRequest(w, common)
		return
	}

	// save file.
	saveFilename := h.GetRandomTempFileName("upload_", fh.Filename)
	h.SaveToFile(r, "file", saveFilename)

	// request visenze/dicoversearch API.
	discoversearch, err := h.aUsecase.GetDiscoversearch(request.Page, request.ResultLimit, saveFilename)
	if err != nil {
		h.Logger.WithFields(logrus.Fields{
			"error": err,
		}).Fatal("UsecaseInterface.GetDiscoversearch() error.")
		h.StatusServerError(w)
		return
	}

	common := CommonResponse{Result: 0, Message: ""}
	jsonData := PostVisenzeDiscoversearchResponse{CommonResponse: common, VisenzeDiscoversearchResponse: *discoversearch}
	h.ResponseJSON(w, jsonData)
}

// NewHTTPHandler responses new HTTPHandler instance.
func NewHTTPHandler(bh *handler.BaseHTTPHandler, s *infrastructure.SQL, c *infrastructure.Cache) *HTTPHandler {
	// article set.
	ar := NewRepository(s.Database, c.Conn)
	au := NewUsecase(s.Database, ar)
	return &HTTPHandler{BaseHTTPHandler: *bh, aUsecase: au}
}
