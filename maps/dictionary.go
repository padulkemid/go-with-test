package maps

type (
	Dictionary    map[string]string
	DictionaryErr string
)

const (
	ErrUnknownKey      = DictionaryErr("could not find the key")
	ErrKeyAlreadyExist = DictionaryErr("key already exists")
	ErrKeyDoesntExist  = DictionaryErr("key not exist")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(key string) (string, error) {
	res, ok := d[key]

	if !ok {
		return "", ErrUnknownKey
	}

	return res, nil
}

func (d Dictionary) Add(k, v string) error {
	_, err := d.Search(k)

	switch err {
	case ErrUnknownKey:
		d[k] = v
	case nil:
		return ErrKeyAlreadyExist
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(k, v string) error {
	_, err := d.Search(k)

	switch err {
	case ErrUnknownKey:
		return ErrKeyDoesntExist
	case nil:
		d[k] = v
	default:
		return err
	}

	return nil
}

func (d Dictionary) Delete(k string) {
	delete(d, k)
}
