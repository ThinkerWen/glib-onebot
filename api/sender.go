package api

import (
	"context"
	"encoding/json"
)

func MarshalJSON(action ActionStruct) ([]byte, error) {
	return json.Marshal(&struct {
		Action string                 `json:"action"`
		Params map[string]interface{} `json:"params"`
	}{
		Action: action.GetAction(),
		Params: action.GetParams(),
	})
}

type IDo interface {
	Do(ctx context.Context) error
}

func (a *ActionStruct) Do(ctx context.Context) error {
	var err error
	var data []byte
	if data, err = MarshalJSON(*a); err != nil {
		return err
	}
	println(string(data))
	return nil
}
