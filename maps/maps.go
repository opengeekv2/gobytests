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

func (d Dictionary) checkWordDoesNotExistsAndDo(word string, do func()) error {
	switch _, err := d.Search(word); err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		do()
		return nil
	default:
		return err
	}
}

func (d Dictionary) Delete(word string) error {
	return d.checkWordDoesNotExistsAndDo(word, func() {
		delete(d, word)
	})
}

func (d Dictionary) Update(word string, newDefinition string) error {
	return d.checkWordDoesNotExistsAndDo(word, func() {
		d[word] = newDefinition
	})
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
