package server

import (
	"testing"

	"the-one/internal/app/server/queue"
	saver2 "the-one/internal/pkg/saver"

	"github.com/stretchr/testify/assert"
)

func TestNewDefaultServer(t *testing.T) {
	type args struct {
		saver saver2.Saver
		port  uint16
	}
	tests := []struct {
		name string
		args args
		want *DefaultServer
	}{
		{
			"Mocked Saver & port 9090",
			args{
				saver2.NewMockedSaver(),
				9090,
			},
			&DefaultServer{
				Port: 9090,
				Q:    queue.New(saver2.NewMockedSaver()),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := NewDefaultServer(tt.args.saver, tt.args.port)

			assert.Equal(t, tt.want.Port, s.Port)
			assert.Equal(t, tt.want.Q.GetSaver(), s.Q.GetSaver())
		})
	}
}
