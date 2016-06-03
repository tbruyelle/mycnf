package mycnf

import (
	"fmt"
	"gopkg.in/gcfg.v1"
	"os"
)

const (
	MyCnfFile = ".my.cnf"
)

type MyCnf struct {
	Client struct {
		User     string
		Password string
	}
}

func Load() (*MyCnf, error) {
	myCnfFile := fmt.Sprintf("%s/%s", os.Getenv("HOME"), MyCnfFile)
	var mycnf MyCnf
	err := gcfg.ReadFileInto(&mycnf, myCnfFile)
	if err != nil {
		return nil, err
	}
	return &mycnf, nil
}
