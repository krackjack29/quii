package dictionary

import "testing"

func TestSearch(t *testing.T) {

	t.Run("known key", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		got, _ := dictionary.Search("test")
		want := "this is just a test"
		assertStrings(t, got, want)
	})

	t.Run("unknown key", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		_, err := dictionary.Search("unknown")
		if err == nil {
			t.Fatal("expected to get an error")
		}
		assertStrings(t, err.Error(), ErrNotFound.Error())
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		value := "this is just a test"
		dictionary.Add("test", value)
		got, err := dictionary.Search("test")
		assertError(t, err, nil)
		assertStrings(t, got, value)
	})

	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{}
		value := "this is just a test"

		dictionary.Add("test", value)

		err := dictionary.Add("test", value)
		assertError(t, err, ErrKeyAlreadyExists)

		got, err := dictionary.Search("test")
		assertError(t, err, nil)
		assertStrings(t, got, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update with adding", func(t *testing.T) {
		dictionary := Dictionary{}
		value := "this is just a test"
		dictionary.Add("test", value)

		dictionary.Update("test", "new value")
		got, err := dictionary.Search("test")
		assertError(t, err, nil)
		assertStrings(t, got, "new value")
	})

	t.Run("update without adding", func(t *testing.T) {
		dictionary := Dictionary{}

		err := dictionary.Update("test", "new value")
		assertError(t, err, ErrNotFound)

	})
}

func TestDelete(t *testing.T) {
	t.Run("delete happy path", func(t *testing.T) {
		key := "test"
		dictionary := Dictionary{key: "what am i doing?"}

		dictionary.Delete(key)

		_, err := dictionary.Search(key)
		assertError(t, err, ErrNotFound)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
