package queue

import "testing"

func TestJosephusProblem(t *testing.T) {
	testCases := []struct{
		name string
		inputs [2]int
		exp int
		err error
	}{
		{
			name: "happy path test case",
			inputs: [2]int{5, 2},
			exp: 3,
			err: nil,
		},
	}
	
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if got, err := JosephusProblem(tc.inputs[0], tc.inputs[1]); got != tc.exp {
				t.Errorf("expected: %d; got: %d", tc.exp, got)
			} else if err != nil && err.Error() != tc.err.Error() {
				t.Errorf("unexpected error occurred: %#v", err)
			}
		})
	}
}