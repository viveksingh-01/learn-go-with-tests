package dictionary

import "errors"

type Dictionary map[string]string

func (d Dictionary) Search(word string) (string, error) {
	// Here, we are using an interesting property of the map lookup.
	// It can return 2 values. The second value is a boolean which indicates if the key was found successfully.
	// This property allows us to differentiate between a word that doesn't exist and a word that
	// just doesn't have a definition.
	definition, ok := d[word]
	if !ok {
		return "", errors.New("could not find the word you were looking for")
	}
	return definition, nil
}
