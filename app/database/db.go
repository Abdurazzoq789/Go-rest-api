package database

import (
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"postapi/app/models"
)

type PostDB interface {
	Open() error
	Close() error
	CreatePost(p *models.Post) error
	GetPosts() ([]*models.Post, error)
	GetPost(r *http.Request) ([]*models.Post, error)
	DeletePost(r *http.Request) error
	UpdatePost(p *models.Post) ([]*models.Post, error)
}

type DB struct {
	db *sqlx.DB
}

func (d *DB) Open() error  {
	pg, err := sqlx.Open("postgres", pgConnStr)
	if err != nil {
		return err
	}

	log.Println("Connected to Database!")

	pg.MustExec(createSchema)

	d.db = pg

	return nil
}

func (d *DB) Close() error  {
	return d.db.Close()
}
