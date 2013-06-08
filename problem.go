package csp

type Value int
type Variable struct {
	Name   string
	Domain []Value
	Value  Value
}

type Solution []Variable
type Constraint func(Solution) bool

type Problem struct {
	Variables   Solution
	Constraints []Constraint
}

func (s Solution) FindVarByName(name string) Variable {
	for _, v := range s {
		if v.Name == name {
			return v
		}
	}

	return Variable{}
}

func (s Solution) GetValues() []Value {
	res := make([]Value, len(s))
	for i, v := range s {
		res[i] = v.Value
	}

	return res
}

// O(n^2)
func (p *Problem) IsDomainConsistent(s Solution) bool {
	for _, v := range s {
		vprime := p.Variables.FindVarByName(v.Name)
		if contains(vprime.Domain, v.Value) != true {
			return false
		}
	}
	return true
}

// O(n)
func (p *Problem) IsDomainConsistentAligned(s Solution) bool {
	for i, v := range s {
		if contains(p.Variables[i].Domain, v.Value) != true {
			return false
		}
	}
	return true
}

func (p *Problem) IsConstraintConsistent(s Solution) bool {
	for _, constraint_func := range p.Constraints {
		if constraint_func(s) != true {
			return false
		}
	}
	return true
}

func (p *Problem) IsConsistent(s Solution) bool {
	if p.IsDomainConsistent(s) != true {
		return false
	}
	if p.IsConstraintConsistent(s) != true {
		return false
	}

	return true
}
