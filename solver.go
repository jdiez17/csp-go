package csp

import (
    "sync"
)

func (p *Problem) doBacktrack(vars Solution, i int, solutions chan<- Solution) {
	if i == len(vars) {
		return
	}

	for _, value := range vars[i].Domain {
		vars[i].Value = value
		p.doBacktrack(vars, i+1, solutions)
		if i == len(vars)-1 {
			// Last node, send a copy of the solution.
            solution := make(Solution, len(vars))
            copy(solution, vars)
			solutions <- solution
		}
	}
}

func solverWorker(p *Problem, input <-chan Solution, output chan<- Solution, wg *sync.WaitGroup) {
    for solution := range input {
        if(p.IsConstraintConsistent(solution)) {
            output <- solution
        }
    }

    wg.Done()
}

func (p *Problem) SolveBacktracking() []Solution {
	solvers := 100 

    var wg sync.WaitGroup

	vars := make(Solution, len(p.Variables))
	copy(vars, p.Variables)

	input := make(chan Solution)
	output := make(chan Solution)

    for i := 0; i < solvers; i++ {
        wg.Add(1)
		go solverWorker(p, input, output, &wg)
	}

    result := make([]Solution, 0) 
    go func() {
        for solution := range output {
            result = append(result, solution)
        }
    }()

	p.doBacktrack(vars, 0, input)
    close(input)

    wg.Wait()
    close(output)
    return result
}
