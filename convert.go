package main

import (
	"context"
	"encoding/json"

	"github.com/go-fed/activity/streams"
	"github.com/go-fed/activity/streams/vocab"
)

func serialize(a vocab.Type) ([]byte, error) {
	var jsonmap map[string]interface{}
	jsonmap, err := streams.Serialize(a)
	if err != nil {
		return nil, err
	}
	b, err := json.Marshal(jsonmap)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func parseGeneric(b []byte) (vocab.Type, error) {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}
	var t vocab.Type
	t, err = streams.ToType(context.Background(), m)
	if err != nil {
		return t, err
	}
	return t, nil
}

func resolve(t vocab.Type, callback interface{}) error {
	typeResolver, err := streams.NewTypeResolver(callback)
	if err != nil {
		return err
	}
	err = typeResolver.Resolve(context.Background(), t)
	if err != nil {
		return err
	}
	return nil
}

func parseSpecific(b []byte, callback interface{}) error {
	var m map[string]interface{}
	err := json.Unmarshal(b, &m)
	if err != nil {
		return err
	}
	resolver, err := streams.NewJSONResolver(callback)
	if err != nil {
		return err
	}
	return resolver.Resolve(context.Background(), m)
}
