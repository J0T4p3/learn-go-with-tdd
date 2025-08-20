package dictionary

import (
	"testing"
)

func TestSearch(t *testing.T) {
	const value = "Testing maps"
	const key = "test"
	dict := Dictionary{key: value}
	t.Run("Search for included items", func(t *testing.T) {
		got, _ := dict.Search(key)
		want := value
		assertString(t, got, want)
	})

	t.Run("Search for not included items", func(t *testing.T) {
		_, got := dict.Search("not a key")
		if got == nil {
			t.Fatal("Expected an error, none was found")
		}
		assertError(t, got, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("add new word", func(t *testing.T) {
		const key = "key"
		const value = "value"
		dict := Dictionary{}
		err := dict.Add(key, value)
		assertError(t, nil, err)
		assertWordOnDict(t, dict, key, value)
	})

	t.Run("add existing word", func(t *testing.T) {
		const key = "key"
		const value = "value"
		dict := Dictionary{key: value}

		err := dict.Add(key, value)
		// error must exist on this case
		if err == nil {
			t.Fatalf("tryed to insert existing word: %v", key)
		}
		assertError(t, err, ErrAlreadyExists)
		assertWordOnDict(t, dict, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update word definition", func(t *testing.T) {
		const key = "key"
		const value = "value"
		const newValue = "new value"
		dict := Dictionary{key: value}
		err := dict.Update(key, newValue)

		assertError(t, nil, err)
		assertWordOnDict(t, dict, key, newValue)
	})

	t.Run("update missing key", func(t *testing.T) {
		const key = "key"
		const value = "value"
		const newValue = "new value"
		dict := Dictionary{key: value}

		err := dict.Update("not a key", newValue)
		assertError(t, err, ErrNotFound)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete word", func(t *testing.T) {
		const key = "key"
		const value = "value"

		dict := Dictionary{key: value}
		err := dict.Delete(key)

		assertError(t, err, nil)
		assertWordNotOnDict(t, dict, key)
	})

	t.Run("delete not existent word", func(t *testing.T) {
		dict := Dictionary{}
		err := dict.Delete("not a key")

		assertError(t, err, ErrNotFound)
	})
}

func assertError(t testing.TB, got, want error) {
	t.Helper()
	if got != want {
		t.Errorf("expected error %q got error %q", got, want)
	}
}

func assertWordNotOnDict(t testing.TB, dict Dictionary, key string) {
	t.Helper()
	_, err := dict.Search(key)
	if err == nil {
		t.Fatalf("shouldn't find searched word: %v", key)
	}
}

func assertWordOnDict(t testing.TB, dict Dictionary, key, value string) {
	t.Helper()
	got, err := dict.Search(key)
	if err != nil {
		t.Fatalf("should find expected word: %v", value)
	}
	assertString(t, got, value)
}

func assertString(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
