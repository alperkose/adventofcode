package input

import "strconv"

type IntContainer struct {
	values []int
}

func (ic *IntContainer) Accept(in string) {
	v, err := strconv.Atoi(in)
	if err != nil {
		return
	}
	ic.values = append(ic.values, v)
}

func (ic *IntContainer) Values() []int {
	return ic.values
}
