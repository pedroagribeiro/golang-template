package optional

import (
	"database/sql/driver"
	"encoding/json"
	"encoding/xml"
	"errors"
	"fmt"
	"strings"

	"gopkg.in/yaml.v3"
)

type optionable interface {
	Scan(any) (any, error)
	Value() (driver.Value, error)
}

type Optional[T optionable] struct {
	value   T
	defined bool
}

func (opt Optional[T]) String() string {
	if opt.Present() {
		return fmt.Sprintf("%v", opt.value)
	}
	return ""
}

func New[DataType optionable](v DataType) Optional[DataType] {
	return Optional[DataType]{value: v, defined: true}
}

func NewUndefined[DataType optionable]() Optional[DataType] {
	return Optional[DataType]{defined: false}
}

func Merge[T optionable](newVal Optional[T], refVal Optional[T]) Optional[T] {
	if newVal.defined {
		return newVal
	}
	return refVal
}

func (opt Optional[T]) Merge(v T) T {
	if opt.Present() {
		return opt.value
	}
	return v
}

func IsEqual[T optionable](val1 Optional[T], val2 Optional[T]) bool {
	if !val1.defined && !val2.defined {
		return true
	}
	if val1.defined != val2.defined {
		return false
	}
	v1, err1 := val1.value.Value()
	v2, err2 := val2.value.Value()

	if err1 != nil || err2 != nil {
		return false
	}
	return v1 == v2
}

func (opt Optional[T]) IsEqual(v T) bool {
	if opt.Present() {
		v1, err1 := opt.value.Value()
		v2, err2 := v.Value()

		if err1 != nil || err2 != nil {
			return false
		}
		return v1 == v2
	}
	return false
}

func (opt *Optional[T]) Set(v T) {
	opt.value = v
	opt.defined = true
}

func (opt Optional[T]) Get() (T, error) {
	if !opt.Present() {
		var nilOpt Optional[T]
		return nilOpt.value, errors.New("value not present")
	}
	return opt.value, nil
}

func (opt Optional[T]) Present() bool {
	return opt.defined
}

func (opt Optional[T]) OrElse(v T) T {
	if opt.Present() {
		return opt.value
	}
	return v
}

func (opt Optional[T]) OrZero() T {
	if opt.Present() {
		return opt.value
	}
	var zero T
	return zero
}

func (opt Optional[T]) Fn(fn func(T)) {
	if opt.Present() {
		fn(opt.value)
	}
}

func (opt Optional[T]) MarshalJSON() ([]byte, error) {
	if opt.Present() {
		return json.Marshal(opt.value)
	}
	return json.Marshal(nil)
}

func (opt *Optional[T]) UnmarshalJSON(data []byte) error {

	if string(data) == "null" {
		(*opt).defined = false
		return nil
	}

	var value T

	if err := json.Unmarshal(data, &value); err != nil {
		// 0x22 is the ASCII code for `"`
		// If the first conversion failed then try again without the surrounding `"`
		if data[0] == 0x22 {
			int_str := string(data)
			int_str = strings.ReplaceAll(int_str, `"`, ``)
			err = json.Unmarshal([]byte(int_str), &value)
			if err != nil {
				return err
			}
		}
	}

	(*opt).value = value
	(*opt).defined = true
	return nil

}

func (opt Optional[T]) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if opt.defined {
		return e.EncodeElement(opt.value, start)
	}
	return nil
}

func (opt *Optional[T]) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var s T
	if err := d.DecodeElement(&s, &start); err != nil {
		return err
	}

	(*opt).value = s
	(*opt).defined = true

	return nil
}

func (opt *Optional[T]) UnmarshalYAML(value *yaml.Node) error {
	// Check if the node is null, which means the value is undefined.
	if value.Kind == yaml.ScalarNode && value.Value == "" {
		opt.defined = false
		return nil
	}

	// Try decoding the value into the generic type T.
	var decodedValue T
	if err := value.Decode(&decodedValue); err != nil {
		return fmt.Errorf("failed to decode YAML into Optional[%T]: %w", decodedValue, err)
	}

	opt.value = decodedValue
	opt.defined = true
	return nil
}

func (opt *Optional[T]) Scan(value interface{}) error {
	if value == nil {
		opt.defined = false
		return nil
	}
	if typedValue, err := opt.value.Scan(value); err != nil {
		return err
	} else {
		if typedValue == nil {
			opt.defined = false
			return nil
		}
		opt.defined, opt.value = true, typedValue.(T)
		return nil
	}
}

func (opt Optional[T]) Value() (driver.Value, error) {
	if !opt.defined {
		return nil, nil
	}
	return opt.value.Value()
}
