package contentful

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewContentfulSDK(t *testing.T) {
	cma, space, err := NewContentfulSDK()

	assert.NoError(t, err, "エラーが発生していないこと")
	assert.NotNil(t, cma, "Contentful SDKのクライアントが生成されていること")
	assert.NotNil(t, space, "Contentful SDKのスペース情報が取得されていること")
}
