package maps

type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("cannot add word because it already exists")
	ErrWordDoesNotExist = DictionaryErr("cannot update word because it does not exist")
)

type Dictionary map[string]string

func (d Dictionary) checkWordDoesNotExists(word string) error {
	switch _, err := d.Search(word); err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		return nil
	default:
		return err
	}
}

func (d Dictionary) Delete(word string) error {
	if err := d.checkWordDoesNotExists(word); err != nil {
		return err
	}
	delete(d, word)
	return nil
}

func (d Dictionary) Update(word string, newDefinition string) error {
	if err := d.checkWordDoesNotExists(word); err != nil {
		return err
	}
	d[word] = newDefinition
	return nil
}

func (d Dictionary) Add(word string, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}
