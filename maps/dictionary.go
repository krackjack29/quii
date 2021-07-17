package dictionary

type Dictionary map[string]string

type DictionaryErr string

func (e DictionaryErr) Error() string { return string(e) }

var (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrKeyAlreadyExists = DictionaryErr("Key already exists")
)

func (d Dictionary) Search(key string) (string, error) {
	value, hasKey := d[key]

	if !hasKey {
		return "", ErrNotFound
	}

	return value, nil
}

func (d Dictionary) Add(key string, value string) error {
	_, err := d.Search(key)

	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrKeyAlreadyExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key string, value string) error {

	_, err := d.Search(key)
	if err != nil {
		return err
	}

	d[key] = value
	return nil
}

func (d Dictionary) Delete(key string) {
	delete(d, key)
}
