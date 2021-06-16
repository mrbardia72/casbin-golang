package memory

import (
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
	"log"
	"strings"
)

type Adapter struct {
	lines []string
}

func NewAdapter() *Adapter {
	return &Adapter{
		lines: make([]string, 0),
	}
}

func (sa *Adapter) LoadPolicy(model model.Model) error {
	log.Println("LoadPolicy")

	for i := range sa.lines {
		persist.LoadPolicyLine(sa.lines[i], model)
	}

	return nil
}

func (sa *Adapter) SavePolicy(model model.Model) error {
	log.Println("SavePolicy")

	return nil
}

func (sa *Adapter) AddPolicy(sec string, ptype string, rule []string) error {
	log.Println("AddPolicy")

	cols := []string{ptype}
	cols = append(cols, rule...)
	line := strings.Join(cols, ", ")

	sa.lines = append(sa.lines, line)

	return nil
}

func (sa *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	log.Println("RemovePolicy")

	cols := []string{ptype}
	cols = append(cols, rule...)
	line := strings.Join(cols, ", ")

	for i, v:= range sa.lines {
		if v == line {
			sa.lines= append(sa.lines[:i], sa.lines[i+1:]...)
		}
	}

	return nil
}

func (sa *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	log.Println("RemoveFilteredPolicy")

	return nil
}
