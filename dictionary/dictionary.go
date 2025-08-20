package dictionary

const (
	ErrNotFound      = DictionaryErr("word inexistent on dictionary")
	ErrAlreadyExists = DictionaryErr("world already on dictionary")
)

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

type Dictionary map[string]string

func (d Dictionary) Search(key string) (string, error) {
	value, ok := d[key]
	if !ok {
		return "", ErrNotFound
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrAlreadyExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(key, value string) error {
	_, err := d.Search(key)

	// If nil, word exists and can be updated
	if err == nil {
		d[key] = value
		return nil
	}

	return ErrNotFound
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)

	// If nil, word exists and can be deleted
	if err == nil {
		delete(d, key)
		return nil
	}

	return ErrNotFound
}
