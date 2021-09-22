package configuration

import (
	"fmt"
	"io/ioutil"
	"path"

	"the-one/internal/saver"

	"gopkg.in/yaml.v2"
)

type SaverConfig struct {
	Type   string
	Params map[string]string
}

type Configuration struct {
	Version int
	Saver   SaverConfig
}

type Dependencies struct {
	Saver saver.Saver
}

func Load(filepath string) (*Dependencies, error) {
	cfg, err := readFile(filepath)
	if err != nil {
		return nil, err
	}
	return buildDependencies(cfg)
}

func readFile(filepath string) (*Configuration, error) {
	fmt.Println(path.Clean(filepath))
	f, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}
	c := &Configuration{}
	err = yaml.UnmarshalStrict(f, c)
	if err != nil {
		return nil, err
	}
	return c, nil
}

func buildDependencies(cfg *Configuration) (*Dependencies, error) {
	s, err := saver.NewSaver(cfg.Saver.Type, cfg.Saver.Params)
	if err != nil {
		return nil, err
	}
	return &Dependencies{
		Saver: s,
	}, nil
}
