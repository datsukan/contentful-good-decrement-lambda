package goodattr

import (
	"testing"

	contentful "github.com/contentful-labs/contentful-go"
	"github.com/stretchr/testify/assert"
)

func TestGoodsAttr(t *testing.T) {
	tests := []struct {
		name string
		good *contentful.Entry
		want int
	}{
		{
			name: "いいね情報が正しく取得できること",
			good: &contentful.Entry{
				Fields: map[string]interface{}{
					"goods": map[string]interface{}{"ja": 10},
				},
			},
			want: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gc, err := GoodsAttr(tt.good)

			assert.NoError(t, err)
			assert.Equal(t, tt.want, gc)
		})
	}
}
