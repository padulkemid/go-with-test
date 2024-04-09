package maps

import "testing"

const (
	mapTestErrorMessage = "got %q want %q given, key is %q"
	mapKeyErrorMessage  = "got %q want %q given"
)

func assertMap(t testing.TB, got, want, key string) {
	t.Helper()

	if got != want {
		t.Errorf(mapTestErrorMessage, got, want, key)
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf(mapKeyErrorMessage, got, want)
	}
}

func assertDefinition(t testing.TB, d Dictionary, key, value string) {
	t.Helper()

	val, err := d.Search(key)
	if err != nil {
		t.Fatal("should have the val: ", err)
	}

	assertMap(t, val, value, key)
}

func TestSearch(t *testing.T) {
	dict := Dictionary{
		"test": "this is just a test",
	}

	t.Run("should get the key because its available", func(t *testing.T) {
		got, _ := dict.Search("test")
		want := "this is just a test"

		assertMap(t, got, want, "test")
	})

	t.Run("shouldn't get the key and throw an err", func(t *testing.T) {
		_, err := dict.Search("unknown")
		want := "could not find the key"

		if err == nil {
			t.Fatal("should throw an error")
		}

		assertMap(t, err.Error(), want, "unknown")
	})
}

func TestAdd(t *testing.T) {
	dict := Dictionary{}

	t.Run("add new dictionary key", func(t *testing.T) {
		err := dict.Add("new", "this is a new key")

		assertError(t, err, nil)
		assertDefinition(t, dict, "new", "this is a new key")
	})

	t.Run("key already exist", func(t *testing.T) {
		err := dict.Add("new", "this is a new key")

		assertError(t, err, ErrKeyAlreadyExist)
		assertDefinition(t, dict, "new", "this is a new key")
	})
}

func TestUpdate(t *testing.T) {
	t.Run("it should update the dict", func(t *testing.T) {
		key := "roar"
		value := "rawr xd"

		dict := Dictionary{
			key: value,
		}

    newVal := "xd aj"

    err := dict.Update(key, newVal)
    
    assertError(t, err, nil)
    assertDefinition(t, dict, key, newVal)
	})

  t.Run("it should check the key if not exist", func(t *testing.T) {
    key := "roar"
    value := "rawr"
    dict := Dictionary{}

    err := dict.Update(key, value)

    assertError(t, err, ErrKeyDoesntExist)
  })
}

func TestDelete(t *testing.T)  {
  t.Run("should delete the key", func(t *testing.T) {
    key := "rawr"
    val := "roar"
    dict := Dictionary{
      key: val,
    }

    dict.Delete(key)

    _, err := dict.Search(key)

    if err != ErrUnknownKey {
      t.Errorf("Expect %q to be deleted", key)
    }
  })
}
