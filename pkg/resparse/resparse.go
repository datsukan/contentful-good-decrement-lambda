package resparse

import (
	"encoding/json"
	"fmt"

	"github.com/datsukan/contentful-good-decrement-lambda/response"
)

// FieldToInt はContentfulのFieldを数値に変換する
func FieldToInt(field interface{}) (int, error) {
	byte, err := json.Marshal(field)
	if err != nil {
		fmt.Println(err)
		return 0, err
	}

	var body response.LangNum
	if err := json.Unmarshal(byte, &body); err != nil {
		fmt.Println(err)
		return 0, err
	}

	return int(body.Ja), nil
}

// IntToField は数値をContentfulのFieldに変換する
func IntToField(num int) (interface{}, error) {
	byte, err := json.Marshal(&response.LangNum{Ja: float64(num)})
	if err != nil {
		fmt.Println(err)
		return "", err
	}

	var body map[string]interface{}
	if err := json.Unmarshal(byte, &body); err != nil {
		fmt.Println(err)
		return "", err
	}

	return body, nil
}
