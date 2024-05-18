package base

import "reflect"

type responseHttp struct {
	Data   data   `json:"data"`
	Errors string `json:"errors,omitempty"`
}

type data struct {
	Records interface{} `json:"records,omitempty"`
	Record  interface{} `json:"record,omitempty"`
}

func SetHttpResponse(result interface{}, err string) interface{} {
	dt := data{}

	isSlice := false
	reflectResult := reflect.ValueOf(result)
	if reflectResult.Kind() == reflect.Ptr {
		isSlice = reflectResult.Elem().Kind() == reflect.Slice
	} else {
		isSlice = reflectResult.Kind() == reflect.Slice
	}

	if isSlice {
		dt.Records = result
		dt.Record = nil
	} else {
		dt.Records = nil
		dt.Record = result
	}

	return responseHttp{
		Data:   dt,
		Errors: err,
	}
}

func GetHttpResponse(resp interface{}) *responseHttp {
	result, ok := resp.(responseHttp)

	if ok {
		return &result
	}

	return nil
}
