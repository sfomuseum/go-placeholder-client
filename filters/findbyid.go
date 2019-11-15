package filters

import (
	"errors"
	"fmt"
	"strings"
)

type FindByIdFilter struct {
	Filter
	key   string
	value string
}

func (f *FindByIdFilter) Key() string {
	return f.key
}

func (f *FindByIdFilter) Value() string {
	return f.value
}

type FindByIdFilters []Filter

func (f *FindByIdFilters) String() string {
	return fmt.Sprintf("%v", *f)
}

func (f *FindByIdFilters) Set(value string) error {

	value = strings.Trim(value, " ")
	kv := strings.Split(value, "=")

	if len(kv) != 2 {
		return errors.New("Invalid search filter")
	}

	switch kv[0] {
	case "lang":
		// pass
	default:
		return errors.New("Invalid findbyid filter")
	}

	sf := FindByIdFilter{
		key:   kv[0],
		value: kv[1],
	}

	*f = append(*f, &sf)
	return nil
}
