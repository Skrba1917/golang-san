package main

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"io"
	"net/http"
)

func decodeBody(r io.Reader) ([]*Config, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var rt []*Config
	if err := dec.Decode(&rt); err != nil {
		return nil, err
	}
	return rt, nil
}

func decodeConfigurationBody(r io.Reader) (*Config, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var rt *Config
	if err := dec.Decode(&rt); err != nil {
		return nil, err
	}
	return rt, nil
}

func decodeGroupBody(r io.Reader) (*Group, error) {
	dec := json.NewDecoder(r)
	dec.DisallowUnknownFields()

	var rt *Group
	if err := dec.Decode(&rt); err != nil {
		return nil, err
	}
	return rt, nil
}

func renderJSON(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func createId() string {
	return uuid.New().String()
}

const (
	configs             = "configs/%s/%s/%s"
	groups              = "groups/%s/%s"
	configurationDelete = "configs/%s/%s"
)

func generateKeyConfigurations(verzija string, labela string) (string, string) {
	id := uuid.New().String()
	return fmt.Sprintf(configs, id, verzija, labela), id
}

func constructKeyConfigurations(id string, verzija string, labela string) string {
	return fmt.Sprintf(configs, id, verzija, labela)
}

func generateKeyGroup(verzija string) (string, string) {
	id := uuid.New().String()
	return fmt.Sprintf(groups, id, verzija), id
}

func constructKeyGroup(id string, verzija string) string {
	return fmt.Sprintf(groups, id, verzija)
}

func constructKeyConfigDelete(id string, Version string) string {
	return fmt.Sprintf(configurationDelete, id, Version)
}
