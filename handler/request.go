package handler

import "fmt"

func errParamIsRequired(name string, typ string) error {
	return fmt.Errorf("param: %s (type: %s) is required", name, typ)
}

type SetContractValueRequest struct {
	Value string `json:"value"`
}

func (r *SetContractValueRequest) Validate() error {
	if r.Value == "" {
		return errParamIsRequired("value", "string")
	}
	return nil
}
