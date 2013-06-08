package csp

import (
	"fmt"
	"testing"
)

// From XKCD: http://xkcd.com/287/
func getSolveableProblem() Problem {
	domain := []Value{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	problem := Problem{
		Variables: []Variable{
			Variable{Name: "fruit", Domain: domain},
			Variable{Name: "frenchfries", Domain: domain},
			Variable{Name: "salad", Domain: domain},
			Variable{Name: "wings", Domain: domain},
			Variable{Name: "sticks", Domain: domain},
			Variable{Name: "sampler", Domain: domain},
		},
		Constraints: []Constraint{
			func(s Solution) bool {
				fruit := s.FindVarByName("fruit").Value
				frenchfries := s.FindVarByName("frenchfries").Value
				salad := s.FindVarByName("salad").Value
				wings := s.FindVarByName("wings").Value
				sticks := s.FindVarByName("sticks").Value
				sampler := s.FindVarByName("sampler").Value

				return fruit*215+frenchfries*275+salad*335+wings*355+sticks*420+sampler*580 == 1505
			},
		},
	}

	return problem
}

func Test_SolveBacktracking(t *testing.T) {
	problem := getSolveableProblem()
	solution := problem.SolveBacktracking()

	if len(solution) == 0 {
		t.Fatal("No solution found")
	}

	for _, sol := range solution {
		fmt.Println(sol.GetValues())
		if problem.IsConsistent(sol) != true {
			t.Error("Solution found by backtracking is inconsistent", sol)
		}
	}
}
