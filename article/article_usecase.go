package article

import (
	"github.com/jinzhu/gorm"
	"goweb1/shared/usecase"
	"github.com/pkg/errors"
)

// UsecaseInterface interface.
type UsecaseInterface interface {
	Get(string) ([]Response, error)
	GetID(int) ([]Response, error)
	Add(string, string) (int64, error)
	Delete(int) (int64, error)
	GetCount() (int, error)
	AddCount() (int, error)
	GetDiscoversearch(int, int, string) (*VisenzeDiscoversearchResponse, error)
}

// Usecase struct.
type Usecase struct {
	usecase.BaseUsecase
	db          *gorm.DB
	aRepository RepositoryInterface
}

// Get search by title.
func (u *Usecase) Get(title string) ([]Response, error) {
	var article []Article
	var err error
	if title == "" {
		article, err = u.aRepository.FindAll()
		if err != nil {
			errors.Wrap(err, "RepositoryInterface.FindAll() error.")
		}
	} else {
		article, err = u.aRepository.Find(title)
		if err != nil {
			errors.Wrap(err, "RepositoryInterface.Find() error.")
		}
	}
	responseArticle := []Response{}

	for _, v := range article {
		responseArticle = append(responseArticle, Response{ID: v.ID, Title: v.Title, Content: v.Content})
	}
	return responseArticle, err
}

// GetID search by id.
func (u *Usecase) GetID(id int) ([]Response, error) {
	article, err := u.aRepository.FindByID(id)
	if err != nil {
		errors.Wrap(err, "RepositoryInterface.FindByID() error.")
	}
	responseArticle := []Response{}

	for _, v := range article {
		responseArticle = append(responseArticle, Response{ID: v.ID, Title: v.Title, Content: v.Content})
	}
	return responseArticle, err
}

// Add method add or update data.
func (u *Usecase) Add(title, content string) (int64, error) {
	var err error
	articles, err := u.Get(title)
	if err != nil {
		errors.Wrap(err, "RepositoryInterface.Get() error.")
		return 0, err
	}

	if len(articles) > 0 {
		return u.aRepository.Update(title, content)
	}
	rowAffected, err := u.aRepository.Create(title, content)
	if err != nil {
		errors.Wrap(err, "RepositoryInterface.Create() error.")
	}
	return rowAffected, err
}

// Delete method delete data.
func (u *Usecase) Delete(id int) (int64, error) {
	// example: transation
	tx := u.db.Begin()
	rowAffected, err := u.aRepository.Delete(tx, id)
	if err != nil {
		errors.Wrap(err, "RepositoryInterface.Delete() error.")
		tx.Rollback()
	} else {
		tx.Commit()
	}
	return rowAffected, err
}

// GetCount get redis counter.
func (u *Usecase) GetCount() (int, error) {
	count, err := u.aRepository.GetCount()
	if err != nil {
		errors.Wrap(err, "RepositoryInterface.GetCount() errors.")
	}
	return count, err
}

// AddCount get redis counter.
func (u *Usecase) AddCount() (int, error) {
	count, err := u.aRepository.AddCount()
	if err != nil {
		errors.Wrap(err, "RepositoryInterface.AddCount() error.")
	}
	return count, err
}

// GetDiscoversearch get result from visenze/discoversearch API.
// example: http://developers.visenze.com/api/#multiple-product-search
func (u *Usecase) GetDiscoversearch(page int, resultLimits int, uploadFile string) (*VisenzeDiscoversearchResponse, error) {
	discoversearch, err := u.aRepository.GetDiscoversearch(page, resultLimits, uploadFile)
	if err != nil {
		errors.Wrap(err, "RepositoryInterface.GetDiscoversearch() errors")
	}
	response := &VisenzeDiscoversearchResponse{}

	// set status.
	response.Status = discoversearch.Status

	// set objects.
	for i, v := range discoversearch.Objects {
		response.Objects = append(response.Objects, ObjectResponse{})
		response.Objects[i].Type = v.Type
		response.Objects[i].Attributes = AttributesResponse{}
		response.Objects[i].Box = v.Box
		for _, vo2 := range v.Result {
			valueMapResponse := ValueMapResponse{
				IMURL: vo2.ValueMap.IMURL,
			}
			resultResponse := ResultResponse{
				IMName:   vo2.IMName,
				Score:    vo2.Score,
				ValueMap: valueMapResponse,
			}
			response.Objects[i].Result = append(response.Objects[i].Result, resultResponse)

		}
		response.Objects[i].Score = v.Score
		response.Objects[i].Total = v.Total
	}

	// set object_types_list
	for il, vl := range discoversearch.ObjectTypesList {
		response.ObjectTypesList = append(response.ObjectTypesList, ObjectTypesListResponse{})
		response.ObjectTypesList[il].Type = vl.Type
		response.ObjectTypesList[il].AttributesList = AttributesResponse{}
	}

	return response, err
}

// NewUsecase responses new Usecase instance.
func NewUsecase(database *gorm.DB, r RepositoryInterface) *Usecase {
	return &Usecase{db: database, aRepository: r}
}
