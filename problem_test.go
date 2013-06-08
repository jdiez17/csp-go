package csp

import (
	"testing"
)

func getProblem() Problem {
	vars := []Variable{
		Variable{Name: "X", Domain: []Value{1, 2, 3}},
	}

	constraints := []Constraint{
		func(s Solution) bool {
			return s.FindVarByName("X").Value == 2 || s.FindVarByName("X").Value == 6
		},
	}

	problem := Problem{
		Variables:   vars,
		Constraints: constraints,
	}

	return problem
}

func Test_IsConsistent(t *testing.T) {
	problem := getProblem()

	solution := Solution{
		Variable{Name: "X", Value: 2},
	}
	wrong_solution := Solution{
		Variable{Name: "X", Value: 3},
	}
	wrong_domain_solution := Solution{
		Variable{Name: "X", Value: 6},
	}

	if problem.IsConsistent(solution) != true {
		t.Error("Consistent solution regarded as inconsistent.")
	}
	if problem.IsConsistent(wrong_solution) {
		t.Error("Inconsistent solution regarded as consistent.")
	}
	if problem.IsConsistent(wrong_domain_solution) {
		t.Error("Solution with a wrong domain regarded as consistent.")
	}
}

func Test_IsDomainConsistentAligned(t *testing.T) {
	problem := getProblem()

	solution := Solution{
		Variable{Name: "X", Value: 2},
	}

	wrong_domain := Solution{
		Variable{Name: "X", Value: 6},
	}

	if problem.IsDomainConsistentAligned(solution) != true {
		t.Error("Solution with correct domain regarded as inconsistent")
	}
	if problem.IsDomainConsistentAligned(wrong_domain) {
		t.Error("Solution with wrong domain regardad as consistent.")
	}
}
