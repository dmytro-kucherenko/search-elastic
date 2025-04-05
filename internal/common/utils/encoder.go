package utils

import (
	"time"

	"github.com/go-playground/validator/v10"
	"github.com/mitchellh/mapstructure"
)

func DecodeStruct[T any](data any, schema T) (err error) {
	decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
		Result:           schema,
		WeaklyTypedInput: true,
		DecodeHook:       mapstructure.StringToTimeHookFunc(time.DateTime),
	})

	if err != nil {
		return
	}

	err = decoder.Decode(data)
	if err != nil {
		return
	}

	validate := validator.New()
	err = validate.Struct(schema)
	if err != nil {
		return
	}

	return
}
