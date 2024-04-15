package reflection

import (
	"fmt"
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

func assertContains(t testing.TB, haystack []string, needle string) {
	t.Helper()
	contains := false

	for _, hay := range haystack {
		if hay == needle {
			contains = true
		}
	}

	if !contains {
		t.Errorf("expected %v to have %q but it didn't", haystack, needle)
	}
}

func TestWalk(t *testing.T) {
	t.Run("walk the path", func(t *testing.T) {
		var got []string
		want := "Padul"

		x := struct {
			Name string
		}{want}

		walk(x, func(input string) {
			got = append(got, input)
		})

		if len(got) != 1 {
			t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
		}

		if got[0] != want {
			t.Errorf("got %q, want %q", got[0], want)
		}
	})

	t.Run("run the cases", func(t *testing.T) {
		cases := []struct {
			Name          string
			Input         interface{}
			ExpectedCalls []string
		}{
			{
				"struct with one string field",
				struct {
					Name string
				}{"Padul"},
				[]string{"Padul"},
			},
			{
				"struct with two string field",
				struct {
					Name string
					City string
				}{"Padul", "Bogor"},
				[]string{"Padul", "Bogor"},
			},
			{
				"struct with non string field",
				struct {
					Name string
					Age  int
				}{"Padul", 17},
				[]string{"Padul"},
			},
			{
				"struct with nested fields",
				Person{
					"Padul",
					Profile{17, "Bogor"},
				},
				[]string{"Padul", "Bogor"},
			},
			{
				"struct with pointers",
				&Person{
					"Padul",
					Profile{17, "Bogor"},
				},
				[]string{"Padul", "Bogor"},
			},
			{
				"struct as a slices",
				[]Profile{
					{17, "London"},
					{19, "Reykjavik"},
				},
				[]string{"London", "Reykjavik"},
			},
			{
				"struct as an array",
				[2]Profile{
					{11, "Jakarta"},
					{12, "Bogor"},
				},
				[]string{"Jakarta", "Bogor"},
			},
		}

		for _, c := range cases {
			t.Run(c.Name, func(t *testing.T) {
				var got []string

				walk(c.Input, func(in string) {
					got = append(got, in)
				})

				if !reflect.DeepEqual(got, c.ExpectedCalls) {
					t.Errorf("got %v, want %v", got, c.ExpectedCalls)
				}
			})
		}
	})

	t.Run("map struct with unordered values", func(t *testing.T) {
		got := make([]string, 0)
		maps := map[string]string{
			"Cat": "Meow",
			"Dog": "Barf",
		}

		walk(maps, func(in string) {
			got = append(got, in)
		})

		assertContains(t, got, "Meow")
		assertContains(t, got, "Barf")
	})

	t.Run("struct with channels", func(t *testing.T) {
		got := make([]string, 0)
		testChan := make(chan Profile)

		go func() {
			testChan <- Profile{12, "Semarang"}
			testChan <- Profile{11, "Amerika"}
			close(testChan)
		}()

		want := []string{"Semarang", "Amerika"}

		walk(testChan, func(in string) {
			got = append(got, in)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})

	t.Run("funcs", func(t *testing.T) {
		// Given
		got := make([]string, 0)
		want := []string{"Sun", "Moon", "Jupiter"}
		testFunc := func(a string) (Profile, Profile, Profile) {
			fmt.Printf("TEST FUNCS -- the string is called as %q \n", a)
			return Profile{11, "Sun"}, Profile{14, "Moon"}, Profile{18, "Jupiter"}
		}

		// When
		walk(testFunc, func(in string) {
			got = append(got, in)
		})

		// Then
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v want %v", got, want)
		}
	})
}
