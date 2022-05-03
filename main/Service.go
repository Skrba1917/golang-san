package main

import (
	"errors"
	"github.com/gorilla/mux"
	"mime"
	"net/http"
)

type Service struct {
	data map[string][]*Config // izigrava bazu podataka
}

func (ts *Service) createPostHandler(w http.ResponseWriter, req *http.Request) {
	contentType := req.Header.Get("Content-Type")
	mediatype, _, err := mime.ParseMediaType(contentType)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if mediatype != "application/json" {
		err := errors.New("Expect application/json Content-Type")
		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
		return
	}

	rt, err := decodeBody(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := createId()
	ts.data[id] = rt
	renderJSON(w, rt)
	w.Write([]byte(id))
}

func (ts *Service) getAllHandler(w http.ResponseWriter, req *http.Request) {
	var allTasks []*Config
	for _, v := range ts.data {
		allTasks = append(allTasks, v...)
	}

	renderJSON(w, allTasks)
}

func (ts *Service) getPostHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	task, ok := ts.data[id]
	if !ok {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	renderJSON(w, task)
}

func (ts *Service) delPostHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	if v, ok := ts.data[id]; ok {
		delete(ts.data, id)
		renderJSON(w, v)
	} else {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
	}
}

func (ts *Service) updatePostHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]

	rt, err := decodeBody(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ts.data[id] = rt
	renderJSON(w, rt)
	w.Write([]byte(id))

	ts.data[id] = append(ts.data[id], rt[0])

}

func (ts *Service) delAllPostHandler(w http.ResponseWriter, req *http.Request) {
}