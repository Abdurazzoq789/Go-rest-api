package database

import (
	"github.com/gorilla/mux"
	"net/http"
	"postapi/app/models"
)

func (d *DB) CreatePost(p *models.Post) error  {
	res, err := d.db.Exec(insertPostSchema, p.Title, p.Content, p.Author)
	if err != nil {
		return err
	}

	res.LastInsertId()
	return err
}

func (d *DB) GetPosts() ([]*models.Post, error)  {
	var posts []*models.Post
	err := d.db.Select(&posts, "SELECT * FROM post")
	if err != nil {
		return posts, err
	}

	return posts, nil
}

func (d *DB) GetPost(r *http.Request) ([]*models.Post, error)  {
	var post []*models.Post

	id := mux.Vars(r)["id"]

	err := d.db.Select(&post, "SELECT id, title, content, author FROM post WHERE id = $1", id)

	if err != nil {
		return post, err
	}

	return post, nil
}

func (d *DB) DeletePost(r *http.Request) error {
	id := mux.Vars(r)["id"]
	res, err := d.db.Exec(deletePostSchema, id)
	if err != nil {
		return err
	}

	res.RowsAffected()
	return err

}

func (d* DB) UpdatePost(p *models.Post) ([]*models.Post, error)  {
	_, err := d.db.Exec(updatePostSchema, p.Title, p.Content, p.Author, p.ID)

	if err != nil {
		return nil, err
	}

	postID := p.ID

	var post  []*models.Post
	getErr := d.db.Select(&post, "SELECT id, title, content, author FROM post WHERE id = $1", postID)


	if getErr != nil {
		return post, getErr
	}

	return post, err
}
