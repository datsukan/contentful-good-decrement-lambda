package resparse

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFieldToInt(t *testing.T) {
	tests := []struct {
		name  string
		field map[string]interface{}
		want  int
	}{
		{
			name:  "想定通りの値に変換できること 1",
			field: map[string]interface{}{"ja": float64(0)},
			want:  0,
		},
		{
			name:  "想定通りの値に変換できること 2",
			field: map[string]interface{}{"ja": float64(12)},
			want:  12,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := FieldToInt(tt.field)

			assert.NoError(t, err, "エラーが発生していないこと")
			assert.Equal(t, tt.want, result)
		})
	}
}

func TestIntToField(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want map[string]interface{}
	}{
		{
			name: "想定通りの値に変換できること 1",
			num:  0,
			want: map[string]interface{}{"ja": float64(0)},
		},
		{
			name: "想定通りの値に変換できること 2",
			num:  12,
			want: map[string]interface{}{"ja": float64(12)},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := IntToField(tt.num)

			assert.NoError(t, err, "エラーが発生していないこと")
			assert.Equal(t, tt.want, result)
		})
	}
}
