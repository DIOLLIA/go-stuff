package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	type Profile struct {
		Age  int
		City string
	}
	type Person struct {
		Name    string
		Profile Profile
	}

	cases := []struct {
		Name           string
		Input          interface{}
		ExpectedResult []string
	}{
		{
			Name:           "Struct with one string field",
			Input:          struct{ input string }{"Andrei"},
			ExpectedResult: []string{"Andrei"},
		},
		{
			Name:           "Struct with two string fields",
			Input:          struct{ input1, input2 string }{"Andrei", "Lucy"},
			ExpectedResult: []string{"Andrei", "Lucy"},
		},
		{
			Name: "Struct with two different fields but extract string only",
			Input: struct {
				InputStr string
				InputInt int
			}{"Andrei", 11},
			ExpectedResult: []string{"Andrei"},
		},
		{
			Name: "Nested structs",
			Input: Person{
				Name: "Andrei",
				Profile: Profile{
					City: "Jelik",
					Age:  22},
			},
			ExpectedResult: []string{"Andrei", "Jelik"},
		},
		{
			Name: "pointers",
			Input: &Person{
				Name: "Andrei",
				Profile: Profile{
					City: "Spb",
					Age:  33},
			},
			ExpectedResult: []string{"Andrei", "Spb"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var result []string

			walk(test.Input, func(input string) {
				result = append(result, input)
			})
			if !reflect.DeepEqual(test.ExpectedResult, result) {
				t.Errorf("expected %q, got %q", test.ExpectedResult, result)
			}
		})
	}
}
