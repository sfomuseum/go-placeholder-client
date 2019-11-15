package filters

import (
	"errors"
	"fmt"
	"strings"
)

type SearchFilter struct {
	Filter
	key   string
	value string
}

func (f *SearchFilter) Key() string {
	return f.key
}

func (f *SearchFilter) Value() string {
	return f.value
}

type SearchFilters []Filter

func (f *SearchFilters) String() string {
	return fmt.Sprintf("%v", *f)
}

func (f *SearchFilters) Set(value string) error {

	value = strings.Trim(value, " ")
	kv := strings.Split(value, "=")

	if len(kv) != 2 {
		return errors.New("Invalid search filter")
	}

	switch kv[0] {
	case "lang", "placetype", "mode":
		// pass
	default:
		return errors.New("Invalid search filter")
	}

	sf := SearchFilter{
		key:   kv[0],
		value: kv[1],
	}

	*f = append(*f, &sf)
	return nil
}
