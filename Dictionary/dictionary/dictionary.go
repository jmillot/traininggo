package dictionary

import (
	"fmt"
	"time"

	"github.com/dgraph-io/badger"
)

// Dictionary is the first struct
type Dictionary struct {
	db *badger.DB
}

// Entry second type for data
type Entry struct {
	Word       string
	Definition string
	CreatedAt  time.Time
}

func (e Entry) String() string {
	created := e.CreatedAt.Format(time.Stamp)
	return fmt.Sprintf("%-10v\t%-50v%-6v", e.Word, e.Definition, created)
}

// New return a new Dictionary
func New(dir string) (*Dictionary, error) {

	opt := badger.DefaultOptions(dir).WithSyncWrites(true)
	opt.Truncate = true

	db, err := badger.Open(opt)
	if err != nil {
		return nil, err
	}

	dict := &Dictionary{
		db: db,
	}
	return dict, nil
}

// Close ferme la connection à la base de données
func (d *Dictionary) Close() {
	d.db.Close()
}
