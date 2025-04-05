package utils

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"
)

type ScanField struct {
	Title      string
	Field      string
	IsOptional bool
}

func ScanFields[T any](fields []ScanField) (T, error) {
	var params T
	data := make(map[string]any)
	reader := bufio.NewReader(os.Stdin)

	for _, option := range fields {
		fmt.Print(option.Title)
		value, err := reader.ReadString('\n')
		if err != nil {
			return params, err
		}

		value = strings.TrimSuffix(value, "\n")
		if !option.IsOptional || value != "" {
			data[option.Field] = value
		}
	}

	err := DecodeStruct(data, &params)
	if err != nil {
		return params, err
	}

	return params, nil
}

func CreateHandler(ctx context.Context) func(action func(ctx context.Context) (string, error)) {
	return func(action func(ctx context.Context) (string, error)) {
		result, err := action(ctx)
		if err != nil {
			fmt.Printf("\n%v\n", err.Error())

			return
		}

		fmt.Print(result)
	}
}
