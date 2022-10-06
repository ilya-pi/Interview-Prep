
func isNumber(s string) bool {
	/*
				It feels like I can record the possible options in a form of a state machine.

				Basically:

				Integer:

				0 - [+-] -> 1 - [0..9] -> 2T - [0..9] -> 2T
				  - [0..9] -> 2T

				Decimal:

				0 - [+-] -> 1 - [0..9] -> 2T -> [0..9] -> 2T - [.] -> 3T -> [0..9] -> 3T
				         -> 1 - [.] -> 4 - [0..9] -> 3T
					- [0..9] -> 2T
					- [.] -> 4

		    0 - [+-] -> 1 - [0..9] -> 2T -> [0..9] -> 2T - [.] -> 3T -> [0..9] -> 3T
		                  - [.] -> 4 - [0..9] -> 3T
		      - [0..9] -> 2T
		      - [.] -> 4


					2T - [.] -> 3T -> [0..9] -> 3T
					1 - [.] -> 4 - [0..9] -> 3T
					0 - [.] -> 4


				Decimal state machine includes the integer by my mistake already

				All terminating states from Decimal, will have -[eE]->(integer state machine)

	*/

	type State struct {
		edges       map[rune]*State
		terminating bool
	}

	// Let's express the rules in state machine now
	newState := func() *State {
		ans := State{edges: make(map[rune]*State)}
		return &ans
	}

	link0to9 := func(target *State, to *State) {
		for i := '0'; i <= '9'; i++ {
			target.edges[i] = to
		}
	}

	// Integer finite automata
	fa_int_2T := newState()
	fa_int_2T.terminating = true
	link0to9(fa_int_2T, fa_int_2T)
	fa_int_1 := newState()
	link0to9(fa_int_1, fa_int_2T)
	fa_int := newState()
	fa_int.edges['+'] = fa_int_1
	fa_int.edges['-'] = fa_int_1
	link0to9(fa_int, fa_int_2T)

	// Decimal finite automate
	fa_dec_2T := newState()
	fa_dec_2T.terminating = true
	link0to9(fa_dec_2T, fa_dec_2T)
	fa_dec_1 := newState()
	link0to9(fa_dec_1, fa_dec_2T)
	fa_dec := newState()
	fa_dec.edges['+'] = fa_dec_1
	fa_dec.edges['-'] = fa_dec_1
	link0to9(fa_dec, fa_dec_2T)
	// extra decimal transitions
	fa_dec_3T := newState()
	fa_dec_3T.terminating = true
	link0to9(fa_dec_3T, fa_dec_3T)
	fa_dec_2T.edges['.'] = fa_dec_3T
	fa_dec_4 := newState()
	link0to9(fa_dec_4, fa_dec_3T)
	fa_dec_1.edges['.'] = fa_dec_4
	fa_dec.edges['.'] = fa_dec_4
	// add the [eE] transitions into fa_int to all terminating edges
	fa_dec_2T.edges['e'] = fa_int
	fa_dec_2T.edges['E'] = fa_int
	fa_dec_3T.edges['e'] = fa_int
	fa_dec_3T.edges['E'] = fa_int

	cur := fa_dec
	for _, v := range []rune(s) {
		if next, ok := cur.edges[v]; ok {
			cur = next
		} else {
			return false
		}
	}
	return cur.terminating
}
