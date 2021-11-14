package saver

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewSaver(t *testing.T) {
	type args struct {
		saverType string
		params    map[string]string
	}
	tests := []struct {
		name    string
		args    args
		want    Saver
		wantErr bool
	}{
		{
			"Mock",
			args{
				saverType: "mock",
				params:    nil,
			},
			NewMockedSaver(),
			false,
		},
		{
			"Mock with PArams",
			args{
				saverType: "mock",
				params: map[string]string{
					"key1": "value1",
					"key2": "value2",
				},
			},
			NewMockedSaverWithParams(map[string]string{
				"key1": "value1",
				"key2": "value2",
			}),
			false,
		},
		{
			"Not Found",
			args{
				saverType: "unknown",
				params:    nil,
			},
			nil,
			true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSaver(tt.args.saverType, tt.args.params)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
