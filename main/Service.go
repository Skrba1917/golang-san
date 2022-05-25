package main

import (
	"errors"
	"github.com/gorilla/mux"
	"mime"
	"net/http"
)

type Service struct {
	store *ConfigurationStore
}

func (ts *Service) createConfigurationHandler(w http.ResponseWriter, req *http.Request) {
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

	rt, err := decodeConfigurationBody(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := createId()
	ts.store.PostConfig(rt)
	renderJSON(w, rt)
	w.Write([]byte(id))
}

func (ts *Service) createGroupHandler(w http.ResponseWriter, req *http.Request) {
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

	rt, err := decodeGroupBody(req.Body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	id := createId()
	ts.store.PostGroup(rt)
	renderJSON(w, rt)
	w.Write([]byte(id))
}

func (ts *Service) getConfigHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	verzija := mux.Vars(req)["version"]
	labela := mux.Vars(req)["label"]
	task, ok := ts.store.GetConfig(id, verzija, labela)
	if ok != nil {
		err := errors.New("key not found")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}
	renderJSON(w, task)
}

func (ts *Service) getGroupHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	verzija := mux.Vars(req)["version"]
	task, ok := ts.store.GetGroup(id, verzija)
	if ok != nil {
		err := errors.New("Nije pronadjeno")
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	renderJSON(w, task)
}

func (ts *Service) delConfigurationHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	verzija := mux.Vars(req)["version"]

	konf, err := ts.store.DeleteConfig(id, verzija)

	renderJSON(w, konf)

	if err != nil {
		errors.New("konfiguracija nije pronadjena, stim nije obrisana!")
		http.Error(w, err.Error(), http.StatusNotFound)
	}

}

func (ts *Service) delGroupHandler(w http.ResponseWriter, req *http.Request) {
	id := mux.Vars(req)["id"]
	verzija := mux.Vars(req)["version"]

	konf, err := ts.store.DeleteGroup(id, verzija)

	renderJSON(w, konf)

	if err != nil {
		errors.New("Grupa nije pronadjena, stim nije obrisana!")
		http.Error(w, err.Error(), http.StatusNotFound)
	}

}

//func (ts *Service) updatePostHandler(w http.ResponseWriter, req *http.Request) {
//
//	contentType := req.Header.Get("Content-Type")
//	mediatype, _, err := mime.ParseMediaType(contentType)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	if mediatype != "application/json" {
//		err := errors.New("Expect application/json Content-Type")
//		http.Error(w, err.Error(), http.StatusUnsupportedMediaType)
//		return
//	}
//
//	rt, err := decodeBody(req.Body)
//	if err != nil {
//		http.Error(w, err.Error(), http.StatusBadRequest)
//		return
//	}
//
//	id := mux.Vars(req)["id"]
//
//	zadat, err3 := ts.data[id]
//	if !err3 {
//		err3 := errors.New("Kljuc nije pronadjen")
//		http.Error(w, err3.Error(), http.StatusNotFound)
//		return
//	}
//
//	for _, config := range rt {
//		zadat = append(zadat, config)
//	}
//
//	ts.data[id] = zadat
//	renderJSON(w, zadat)
//
//}

//func (ts *Service) delAllPostHandler(w http.ResponseWriter, req *http.Request) {
//}
