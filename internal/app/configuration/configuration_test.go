package configuration

import (
	"testing"

	"the-one/internal/pkg/saver"

	"github.com/stretchr/testify/assert"
)

func Test_readFile(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		want    *Configuration
		wantErr bool
	}{
		{
			name: "File OK",
			args: args{
				filepath: "test/config.yml",
			},
			want: &Configuration{
				Saver: SaverConfig{
					Type: "mock",
				},
			},
		},
		{
			name: "Full File",
			args: args{
				filepath: "test/config_full.yml",
			},
			want: &Configuration{
				Version: 1,
				Saver: SaverConfig{
					"mock",
					map[string]string{
						"key": "value",
					},
				},
			},
		},
		{
			name: "File Not found",
			args: args{
				filepath: "test/not_found_file",
			},
			wantErr: true,
		},
		{
			name: "Invalid format",
			args: args{
				filepath: "test/config.json",
			},
			wantErr: true,
		},
		{
			name: "Invalid params",
			args: args{
				filepath: "test/invalid.yml",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readFile(tt.args.filepath)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func Test_buildDependencies(t *testing.T) {
	type args struct {
		cfg *Configuration
	}
	tests := []struct {
		name      string
		args      args
		want      *Dependencies
		wantedLog string
		wantErr   bool
	}{
		{
			"OK",
			args{
				cfg: &Configuration{
					Saver: SaverConfig{
						Type: "mock",
						Params: map[string]string{
							"key": "value",
						},
					},
				},
			},
			&Dependencies{
				Saver: saver.NewMockedSaverWithParams(map[string]string{
					"key": "value",
				}),
			},
			"/log/test.log",
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := buildDependencies(tt.args.cfg)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}

func TestLoad(t *testing.T) {
	type args struct {
		filepath string
	}
	tests := []struct {
		name    string
		args    args
		want    *Dependencies
		wantErr bool
	}{
		{
			name: "File OK",
			args: args{
				filepath: "test/config.yml",
			},
			want: &Dependencies{
				Saver: saver.NewMockedSaver(),
			},
		},
		{
			name: "Error reading file",
			args: args{
				filepath: "test/error.yml",
			},
			wantErr: true,
		},
		{
			name: "Error building dependency",
			args: args{
				filepath: "test/invalid_value.yml",
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Load(tt.args.filepath)
			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.want, got)
			}
		})
	}
}
