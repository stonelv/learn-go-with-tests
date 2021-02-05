package reflection

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	expected := "Lyle"
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

	if got[0] != expected {
		t.Errorf("got %q, want %q", got[0], expected)
	}
}

func TestWalkWithCases(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Lyle"},
			[]string{"Lyle"},
		},
		{
			"Struct with two string fields",
			struct {
				Name string
				City string
			}{"Lyle", "Beijing"},
			[]string{"Lyle", "Beijing"},
		},
		{
			"Struct with non string field",
			struct {
				Name string
				Age  int
			}{"Lyle", 38},
			[]string{"Lyle"},
		},
		{
			"Nested fields",
			Person{
				"Lyle",
				Profile{38, "Beijing"},
			},
			[]string{"Lyle", "Beijing"},
		},
		{
			"Pointers to things",
			&Person{
				"Lyle",
				Profile{38, "Beijing"},
			},
			[]string{"Lyle", "Beijing"},
		},
		{
			"Slices",
			[]Profile{
				{33, "London"},
				{38, "Beijing"},
			},
			[]string{"London", "Beijing"},
		},
		{
			"Arrays",
			[2]Profile{
				{33, "Shanghai"},
				{38, "Beijing"},
			},
			[]string{"Shanghai", "Beijing"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}

	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"name": "Lyle",
			"city": "Beijing",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Lyle")
		assertContains(t, got, "Beijing")
	})

	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{35, "Katowice"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{38, "Beijing"}
		}

		var got []string
		want := []string{"Berlin", "Beijing"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func assertContains(t testing.TB, hayStack []string, needle string) {
	t.Helper()
	contains := false
	for _, x := range hayStack {
		if x == needle {
			contains = true
		}
	}
	if !contains {
		t.Errorf("expected %+v to contain %q, but it didn't", hayStack, needle)
	}
}
