package reflection

import (
	"reflect"
	"testing"
)

type Profile struct {
	Age  int
	City string
}
type Person struct {
	Name    string
	Profile Profile
}

func TestWalk(t *testing.T) {

	cases := []struct {
		Name           string
		Input          interface{}
		ExpectedResult []string
	}{
		{
			Name:           "Struct with one string field",
			Input:          struct{ Input string }{"Andrei"},
			ExpectedResult: []string{"Andrei"},
		},
		{
			Name:           "Struct with two string fields",
			Input:          struct{ Input1, Input2 string }{"Andrei", "Lucy"},
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
		{
			Name: "slices",
			Input: []Profile{
				{
					City: "Spb",
					Age:  78},
				{
					City: "Krsk",
					Age:  24},
			},
			ExpectedResult: []string{"Spb", "Krsk"},
		},
		{
			Name: "arrays",
			Input: [2]Profile{
				{
					City: "Spb",
					Age:  78},
				{
					City: "Krsk",
					Age:  24},
			},
			ExpectedResult: []string{"Spb", "Krsk"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var result []string

			walk(test.Input, func(Input string) {
				result = append(result, Input)
			})
			if !reflect.DeepEqual(test.ExpectedResult, result) {
				t.Errorf("expected %q, got %q", test.ExpectedResult, result)
			}
		})
	}
	t.Run("with maps", func(t *testing.T) {
		cities := map[string]string{
			"City":    "Spb",
			"Town":    "Krsk",
			"Village": "Jelik"}

		var result []string
		walk(cities, func(input string) {
			result = append(result, input)
		})
		assertContains(t, "Spb", result)
		assertContains(t, "Krsk", result)
		assertContains(t, "Jelik", result)
	})
}

func TestWalkWithChannels(t *testing.T) {
	walkChan := make(chan Profile)
	go func() {
		walkChan <- Profile{22, "Jelik"}
		walkChan <- Profile{24, "Krsk"}
		close(walkChan)
	}()

	expected := []string{"Jelik", "Krsk"}
	var result []string

	walk(walkChan, func(input string) {
		result = append(result, input)
	})
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %q, got %q", expected, result)
	}
}

func TestWalkFunc(t *testing.T) {
	walkFunc := func() (Profile, Profile) {
		return Profile{22, "Jelik"}, Profile{24, "Krsk"}
	}

	expected := []string{"Jelik", "Krsk"}
	var result []string

	walk(walkFunc, func(input string) {
		result = append(result, input)
	})
	if !reflect.DeepEqual(expected, result) {
		t.Errorf("expected %q, got result %q", expected, result)
	}
}

func assertContains(t testing.TB, expected string, result []string) {
	t.Helper()
	contains := false
	for _, x := range result {
		if x == expected {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %v to contain %q but it didn't", expected, result)
	}
}

func TestWalkInitial(t *testing.T) {

	expected := "Chris"
	var got []string

	x := struct {
		Name string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if len(got) != 1 {
		t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
	}
}
