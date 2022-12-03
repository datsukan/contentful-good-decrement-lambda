package config

import (
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadContentfulEnv(t *testing.T) {
	os.Setenv("CONTENTFUL_ACCESS_TOKEN", "test_access_token")
	os.Setenv("CONTENTFUL_SPACE_ID", "test_space_id")

	wantAccessToken := "test_access_token"
	wantSpaceID := "test_space_id"

	token, spaceID, err := LoadContentfulEnv()

	assert.NoError(t, err, "Contentfulの環境変数の読み込みでエラーが発生しないこと")
	assert.Equal(t, wantAccessToken, token, "Contentfulのアクセストークンの値が正しく読み込まれること")
	assert.Equal(t, wantSpaceID, spaceID, "ContentfulのスペースIDの値が正しく読み込まれること")
}
