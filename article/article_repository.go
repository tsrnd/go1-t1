package article

import (
	"encoding/json"
	"strconv"

	"goweb1/infrastructure"
	"goweb1/shared/repository"

	"github.com/garyburd/redigo/redis"
	"github.com/jinzhu/gorm"
	"github.com/levigross/grequests"
)

// RepositoryInterface interface.
type RepositoryInterface interface {
	Find(string) ([]Article, error)
	FindByID(int) ([]Article, error)
	FindAll() ([]Article, error)
	Create(string, string) (int64, error)
	Update(string, string) (int64, error)
	Delete(*gorm.DB, int) (int64, error)
	GetCount() (int, error)
	AddCount() (int, error)
	GetDiscoversearch(int, int, string) (*VisenzeDiscoversearch, error)
}

// Repository struct.
type Repository struct {
	repository.BaseRepository
	// db connect database.
	db *gorm.DB
	// radis connect redis.
	redis *redis.Conn
}

// Find function select from article table.
func (r *Repository) Find(title string) ([]Article, error) {
	article := []Article{}
	err := r.db.Where("title = ?", title).Find(&article).Error
	return article, err
}

// FindAll function select all from article table.
func (r *Repository) FindAll() ([]Article, error) {
	article := []Article{}
	err := r.db.Find(&article).Error
	return article, err
}

// FindByID function select ID from article table.
func (r *Repository) FindByID(id int) ([]Article, error) {
	article := []Article{}
	err := r.db.Where("id = ?", id).Find(&article).Error
	return article, err
}

// Create function insert article table.
// http://jinzhu.me/gorm/crud.html#create
func (r *Repository) Create(title, content string) (int64, error) {
	article := &Article{Title: title, Content: content}
	create := r.db.Create(article)
	return create.RowsAffected, create.Error
}

// Update function update article table.
// http://jinzhu.me/gorm/crud.html#update
func (r *Repository) Update(title, content string) (int64, error) {
	article := &Article{Title: title, Content: content}
	update := r.db.Model(&Article{}).Where("title = ?", title).Update(article)
	return update.RowsAffected, update.Error
	// UPDATE article SET title='title', content='content' updated_at='2013-11-17 21:34:10' WHERE title='title;
}

// Delete function delete article table.
// http://jinzhu.me/gorm/crud.html#delete
func (r *Repository) Delete(tx *gorm.DB, id int) (int64, error) {
	delete := tx.Where("id = ?", id).Delete(&Article{})
	return delete.RowsAffected, delete.Error
	// UPDATE article SET deleted_at="2013-10-29 10:23" WHERE id = 20;
}

// GetCount method get redis count(key="article").
// https://github.com/garyburd/redigo
// https://redis.io/commands
func (r *Repository) GetCount() (int, error) {
	conn := *r.redis
	count, err := conn.Do("GET", "article")
	if count == nil {
		return 0, err
	}
	return redis.Int(count, err)
}

// AddCount method increment redis count(key="article").
// https://github.com/garyburd/redigo
// https://redis.io/commands
func (r *Repository) AddCount() (int, error) {
	conn := *r.redis
	count, err := conn.Do("INCR", "article")
	return redis.Int(count, err)
}

// GetDiscoversearch method get result from visenze/discoversearch API.
func (r *Repository) GetDiscoversearch(page int, resultLimit int, uploadFile string) (*VisenzeDiscoversearch, error) {
	// get visenze config .
	accessKey := infrastructure.GetConfigString("api.visenze.accesskey")
	secretKey := infrastructure.GetConfigString("api.visenze.secretkey")

	// file upload.
	fd, err := grequests.FileUploadFromDisk(uploadFile)
	if err != nil {
		return nil, err
	}
	fd[0].FieldName = "image"

	data := map[string]string{}
	data["detection_sensitivity"] = "low"
	if page > 0 {
		data["page"] = strconv.Itoa(page)
	}
	if resultLimit > 0 {
		data["result_limit"] = strconv.Itoa(resultLimit)
	}
	data["fl"] = "im_url"
	data["dedup"] = "false"

	// options
	ro := &grequests.RequestOptions{
		Auth:  []string{accessKey, secretKey},
		Files: fd,
		Data:  data,
	}
	resp, err := grequests.Post(URLVisenzeDiscoversearch, ro)

	// response DiscoverSearchResponse struct.
	responseDiscoverSearch := &VisenzeDiscoversearch{}
	//json.Unmarshal(resp.Bytes(), responseDiscoverSearch)
	err = json.NewDecoder(resp).Decode(responseDiscoverSearch)

	return responseDiscoverSearch, err
}

// NewRepository responses new Repository instance.
func NewRepository(database *gorm.DB, redis *redis.Conn) *Repository {
	return &Repository{db: database, redis: redis}
}
