package maps

import "testing"

const key = "key_to_add"
const value = "value_to_add"

func TestSearchByKey(t *testing.T) {

	dictionary := Dictionary{key: value}
	expected := value

	t.Run("get by existing key", func(t *testing.T) {
		result, _ := dictionary.SearchByKey(key)
		assertStrings(t, expected, result)
	})

	t.Run("get by not existing key", func(t *testing.T) {
		_, err := dictionary.SearchByKey("not_existing_key")

		if err == nil {
			t.Fatal("expected got and error")
		}
		assertError(t, ErrorNotFound, err)
	})
}

func assertStrings(t *testing.T, expected, result string) {
	t.Helper()
	if expected != result {
		t.Errorf("Expected value %q but got %q", expected, result)
	}
}

func assertError(t *testing.T, expected, result error) {
	t.Helper()
	if expected != result {
		t.Errorf("Expected error %q, got %q", expected, result)
	}
}
func assertDefinition(t *testing.T, dictionary Dictionary, key, value string) {
	t.Helper()
	result, err := dictionary.SearchByKey(key)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, value, result)
}

func TestAdd(t *testing.T) {
	t.Run("add new key", func(t *testing.T) {
		dictionary := Dictionary{}
		err := dictionary.Add(key, value)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, key, value)
	})
	t.Run("add existing key", func(t *testing.T) {
		dictionary := Dictionary{key: value}

		err := dictionary.Add(key, "new_value")

		assertError(t, ErrorKeyAlreadyExists, err)
		assertDefinition(t, dictionary, key, value)
	})
}

func TestUpdate(t *testing.T) {
	t.Run("update new key", func(t *testing.T) {

		dictionary := Dictionary{key: value}
		newValue := "new_value"
		err := dictionary.Update("new_key", newValue)

		assertError(t, ErrWordDoesNotExist, err)
		assertDefinition(t, dictionary, key, value)
	})

	t.Run("update existing key", func(t *testing.T) {
		dictionary := Dictionary{key: value}
		newValue := "new_value"
		err := dictionary.Update(key, newValue)

		assertError(t, nil, err)
		assertDefinition(t, dictionary, key, newValue)
	})
}

func TestDelete(t *testing.T) {
	t.Run("delete existing key", func(t *testing.T) {
		dictionary := Dictionary{key: value}
		_ = dictionary.Delete(key)

		_, err := dictionary.SearchByKey(key)

		assertError(t, ErrorNotFound, err)
	})
	t.Run("delete non-existent key", func(t *testing.T) {
		dictionary := Dictionary{key: value}
		err := dictionary.Delete("not_existing_key")

		assertError(t, ErrorNotFound, err)
	})
}
