package contentful

import (
	"fmt"

	"github.com/datsukan/contentful-good-decrement-lambda/config"

	contentful "github.com/contentful-labs/contentful-go"
)

// NewContentfulSDK はContentful SDKのクライアントインスタンスを生成する
func NewContentfulSDK() (*contentful.Client, *contentful.Space, error) {
	token, spaceID, err := config.LoadContentfulEnv()
	if err != nil {
		return nil, nil, err
	}

	cma := contentful.NewCMA(token)
	space, err := cma.Spaces.Get(spaceID)

	if err != nil {
		fmt.Println(err)
		return nil, nil, err
	}

	return cma, space, nil
}
