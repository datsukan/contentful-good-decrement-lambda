package goodattr

import (
	"fmt"

	pkgrp "github.com/datsukan/contentful-good-decrement-lambda/pkg/resparse"

	"github.com/contentful-labs/contentful-go"
)

// GoodsAttr は entry からいいね数を取得する
func GoodsAttr(entry *contentful.Entry) (int, error) {
	var count int
	var err error
	for attr, field := range entry.Fields {
		switch attr {
		case "goods":
			count, err = pkgrp.FieldToInt(field)
			if err != nil {
				fmt.Println(err)
				return 0, err
			}
		}
	}

	return count, nil
}
