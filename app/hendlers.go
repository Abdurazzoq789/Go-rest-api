package app

import (
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"postapi/app/models"
	"strconv"
)

func (a *App) IndexHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w, "Welcome to post Api")
		if err != nil {
			log.Println(err)
			os.Exit(1)
		}
	}
}

func (a *App) CreatePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.PostRequest{}
		err := parse(w, r, &req)
		if err != nil {
			log.Printf("Cannot parse post body. err=%v \n", err)
			sendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		p := &models.Post{
			ID:      0,
			Title:   req.Title,
			Content: req.Content,
			Author:  req.Author,
		}

		err = a.DB.CreatePost(p)
		if err != nil {
			log.Printf("Cannot save post in DB. err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		resp := mapPostToJSON(p)
		sendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *App) GetPostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		posts, err := a.DB.GetPosts()
		if err != nil {
			log.Printf("Cannot get posts, err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]models.JsonPost, len(posts))
		for idx, post := range posts {
			resp[idx] = mapPostToJSON(post)
		}

		sendResponse(w, r, resp, http.StatusOK)
	}
}
func (a *App) GetOnePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		post, err := a.DB.GetPost(r)
		if err != nil {
			log.Printf("Cannot get post, err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		if len(post) == 0 {
			log.Printf("The request post doesnt exist")
			data := make(map[string]interface{})
			data["status code"] = -1
			data["message"] = "The request post doesnt exist"
			sendResponse(w, r, data, http.StatusNotFound)
			return
		}

		var resp = make([]models.JsonPost, len(post))
		for idx, postItem := range post {
			resp[idx] = mapPostToJSON(postItem)
		}

		sendResponse(w, r, resp, http.StatusOK)
	}
}

func (a *App) DeletePostHandler() http.HandlerFunc  {
	return func(w http.ResponseWriter, r *http.Request) {
		err := a.DB.DeletePost(r)
		if err != nil {
			log.Printf("Cannot get post, err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		data := map[string]interface{}{
			"code": 1,
			"message" : "Successfully deleted",
		}

		sendResponse(w, r, data, http.StatusOK)
	}
}

func (a *App) UpdatePostHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		req := models.PostRequest{}
		err := parse(w, r, &req)

		if err != nil {
			log.Printf("Cannot parse post body. err=%v \n", err)
			sendResponse(w, r, nil, http.StatusBadRequest)
			return
		}

		postID := mux.Vars(r)["id"]
		id, _ := strconv.ParseInt(postID, 0, 64)

		var p = &models.Post{
			ID:      id,
			Title:   req.Title,
			Content: req.Content,
			Author:  req.Author,
		}
		post, err := a.DB.UpdatePost(p)

		if err != nil {
			log.Printf("Cannot save post in DB. err=%v \n", err)
			sendResponse(w, r, nil, http.StatusInternalServerError)
			return
		}

		var resp = make([]models.JsonPost, 1)
		for idx, postItem := range post {
			resp[idx] = mapPostToJSON(postItem)
		}

		sendResponse(w, r, resp, http.StatusOK)
	}
}

