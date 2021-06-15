package bardiaadapter

import (
	"encoding/json"
	"errors"
	"github.com/casbin/casbin/v2/model"
	"github.com/casbin/casbin/v2/persist"
)

type role struct {
	PType   string
	Subject string
	Role    string
}

type Adapter struct {
	source *[]byte
	roles  []role
}

func NewAdapter(source *[]byte) *Adapter {
	a := Adapter{}
	a.source = source
	a.roles = make([]role, 0)
	return &a
}

func (a *Adapter) saveToMemory() error {
	data, err := json.Marshal(a.roles)
	if err == nil {
		*a.source = data
	}
	return err
}

func (a *Adapter) loadFromMemory() error {
	var policy []role
	err := json.Unmarshal(*a.source, &policy)
	if err == nil {
		a.roles = policy
	}
	return err
}

func (a *Adapter) LoadPolicy(model model.Model) error {

	err := a.loadFromMemory()

	if err != nil {
		return err
	}
	for _, ss := range a.roles {
		loadPolicyLine(ss, model)
	}

	return nil
}

func savePolicyLine(ptype string, rule []string) role {
	line := role{}

	line.PType = ptype

	if len(rule) > 0 {
		line.Subject = rule[0]
	}

	if len(rule) > 1 {
		line.Role = rule[1]
	}

	return line
}

func (a *Adapter) SavePolicy(model model.Model) error {

	a.roles = []role{}

	var lines []role

	for ptype, ast := range model["p"] {
		for _, rule := range ast.Policy {
			line := savePolicyLine(ptype, rule)
			lines = append(lines, line)
		}
	}

	for ptype, ast := range model["g"] {
		for _, rule := range ast.Policy {
			line := savePolicyLine(ptype, rule)
			lines = append(lines, line)
		}
	}

	a.roles = lines

	err := a.saveToMemory()
	return err

}

func loadPolicyLine(line role, model model.Model) {

	lineText := line.PType

	if line.Subject != "" {
		lineText += ", " + line.Subject
	}

	if line.Role != "" {
		lineText += ", " + line.Role
	}

	persist.LoadPolicyLine(lineText, model)

}

func (a *Adapter) AddPolicy(sec string, ptype string, rule []string) error {
	return errors.New("not implemented")
}

func (a *Adapter) RemovePolicy(sec string, ptype string, rule []string) error {
	return errors.New("not implemented")
}

func (a *Adapter) RemoveFilteredPolicy(sec string, ptype string, fieldIndex int, fieldValues ...string) error {
	return errors.New("not implemented")
}
