package id

import (
	"fmt"

	gonanoid "github.com/matoous/go-nanoid/v2"
)

func Generate(t Type) (string, error) {
	id, err := gonanoid.Generate("0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz", ID_LENGTH)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s_%s", t, id), nil
}
