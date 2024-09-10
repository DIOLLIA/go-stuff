package maps

const (
	ErrorNotFound         = DictionaryErr("can't find the value for the given key")
	ErrorKeyAlreadyExists = DictionaryErr("can't add the value that already exist")
	ErrWordDoesNotExist   = DictionaryErr("can't update word because it doesn't exist")
)

type Dictionary map[string]string
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}
func (d Dictionary) Add(key, value string) error {
	_, err := d.SearchByKey(key)
	switch err {
	case ErrorNotFound:
		d[key] = value
	case nil:
		return ErrorKeyAlreadyExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key, newValue string) error {
	if value, _ := d.SearchByKey(key); value == "" {
		return ErrWordDoesNotExist
	}
	d[key] = newValue
	return nil
}
func (d Dictionary) Delete(key string) error {
	if value, _ := d.SearchByKey(key); value == "" {
		return ErrorNotFound
	}
	delete(d, key)
	return nil
}

func (d Dictionary) SearchByKey(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrorNotFound
	}
	return value, nil
}
