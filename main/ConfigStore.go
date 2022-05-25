package main

import (
	"encoding/json"
	"fmt"
	"github.com/hashicorp/consul/api"
	"os"
)

type ConfigurationStore struct {
	cli *api.Client
}

func New() (*ConfigurationStore, error) {
	db := os.Getenv("DB")
	dbport := os.Getenv("DBPORT")

	config := api.DefaultConfig()
	config.Address = fmt.Sprintf("%s:%s", db, dbport)
	client, err := api.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &ConfigurationStore{
		cli: client,
	}, nil
}

func (ps *ConfigurationStore) GetConfig(id string, verzija string, labela string) (*Config, error) {
	kv := ps.cli.KV()

	pair, _, err := kv.Get(constructKeyConfigurations(id, verzija, labela), nil)
	if err != nil {
		return nil, err
	}

	post := &Config{}
	err = json.Unmarshal(pair.Value, post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (ps *ConfigurationStore) GetGroup(id string, verzija string) (*Group, error) {
	kv := ps.cli.KV()

	pair, _, err := kv.Get(constructKeyGroup(id, verzija), nil)
	if err != nil {
		return nil, err
	}

	post := &Group{}
	err = json.Unmarshal(pair.Value, post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (ps *ConfigurationStore) DeleteConfig(id string, verzija string) (map[string]string, error) {
	kv := ps.cli.KV()
	_, err := kv.Delete(constructKeyConfigDelete(id, verzija), nil)
	if err != nil {
		fmt.Println("Konfiguracija nije obrisana!")
		return nil, err

	}

	return map[string]string{"Deleted": id}, nil
}

func (ps *ConfigurationStore) DeleteGroup(id string, verzija string) (map[string]string, error) {
	kv := ps.cli.KV()
	_, err := kv.Delete(constructKeyGroup(id, verzija), nil)
	if err != nil {
		fmt.Println("Grupa nije obrisana!")
		return nil, err
	}

	return map[string]string{"Deleted": id}, nil
}

func (ps *ConfigurationStore) PostConfig(post *Config) (*Config, error) {
	kv := ps.cli.KV()

	sid, rid := generateKeyConfigurations(post.Version, post.Label)
	post.Id = rid

	data, err := json.Marshal(post)
	if err != nil {
		return nil, err
	}

	p := &api.KVPair{Key: sid, Value: data}
	_, err = kv.Put(p, nil)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func (ps *ConfigurationStore) PostGroup(post *Group) (*Group, error) {
	kv := ps.cli.KV()

	sid, rid := generateKeyGroup(post.Version)
	post.Id = rid

	data, err := json.Marshal(post)
	if err != nil {
		return nil, err
	}

	p := &api.KVPair{Key: sid, Value: data}
	_, err = kv.Put(p, nil)
	if err != nil {
		return nil, err
	}

	return post, nil
}
