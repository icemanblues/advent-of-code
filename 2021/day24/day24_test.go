package main

import "testing"

func TestAluInst(t *testing.T) {
	tests := []struct {
		name  string
		model int
		insts []string
		z     int
	}{
		{"add", 23111111111111, []string{"inp w", "inp x", "add z w", "add z x"}, 5},
		{"mul", 23111111111111, []string{"inp w", "inp x", "add z w", "mul z x"}, 6},
		{"div", 82111111111111, []string{"inp z", "inp x", "div z x"}, 4},
		{"div truncate", 83111111111111, []string{"inp z", "inp x", "div z x"}, 2},
		{"mod", 73111111111111, []string{"inp z", "inp x", "mod z x"}, 1},
		{"eql 1", 23111111111111, []string{"inp z", "inp x", "eql z x"}, 0},
		{"eql 0", 33111111111111, []string{"inp z", "inp x", "eql z x"}, 1},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			model, err := IntToModel(test.model)
			if err != nil {
				t.Error(err)
			}

			insts := Instructions(test.insts)
			alu, err := Run(insts, model)
			if err != nil {
				t.Error(err)
			}

			z := alu.registers["z"]
			if z != test.z {
				t.Errorf("expected %v actual %v\n", test.z, z)
			}
		})
	}
}
