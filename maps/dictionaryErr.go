package maps

const (
	//ErrorNotFound used when not found event
	ErrorNotFound = DictionaryErr("could not find the error you were looking for")
	//ErrWordExists used when word already exists
	ErrWordExists = DictionaryErr("cannot add word because it already exists")
	//ErrWordDoesNotExist used when word does not exist
	ErrWordDoesNotExist = DictionaryErr("word does not exist")
)

//DictionaryErr used to store error info
type DictionaryErr string

func (e DictionaryErr) Error() string {
	return string(e)
}