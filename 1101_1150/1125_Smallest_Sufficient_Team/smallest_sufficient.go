package smallestsufficient

type state struct {
	headCount, prePerson, preState int
}

func smallestSufficientTeam(reqSkills []string, people [][]string) []int {
	skillMap := map[string]int{}
	for i, v := range reqSkills {
		skillMap[v] = i
	}
	n := len(reqSkills)
	target := 1<<uint(n) - 1
	states := make([]*state, 1<<uint(n))
	states[0] = &state{}

	// Iterate on people to update states
	for person, skills := range people {
		skillSum := 0

		for _, skill := range skills {
			skillSum |= 1 << uint(skillMap[skill])
		}

		for current := target; current >= 0; current-- {
			// use current person to update sum
			next := skillSum | current

			if next == current {
				// sum is not updated, continue
				continue
			}

			if states[current] != nil {
				if states[next] == nil || ((*states[next]).headCount > (*states[current]).headCount+1) {
					// update next state, if it has not been reached, or can use fewer people.
					if states[next] == nil {
						states[next] = &state{}
					}
					(*states[next]).headCount = (*states[current]).headCount + 1
					(*states[next]).prePerson = person
					(*states[next]).preState = current
				}
			}
		}
	}

	ans := []int{}
	s := target
	for s != 0 {
		ans = append(ans, (*states[s]).prePerson)
		s = (*states[s]).preState
	}
	return ans
}
