package option

import (
	"encoding/json"
	"fmt"
)

type Option[T any] struct {
	value  T
	isFull bool
}

func (s *Option[T]) Get() T {
	if s.isFull {
		return s.value
	} else {
		panic("cannot get from None type")

	}
}

func (s *Option[T]) GetOrElse(other T) T {
	if s.isFull {
		return s.value
	} else {
		return other
	}
}

func (s *Option[T]) IsEmpty() bool {
	return !s.isFull
}

func (s Option[T]) String() string {
	if s.isFull {
		return fmt.Sprintf("Some(%v)", s.value)
	} else {
		return "None"
	}
}

func (s Option[T]) MarshalJSON() ([]byte, error) {
	if s.isFull {
		return json.Marshal(s.value)
	} else {
		return []byte("null"), nil
	}
}

func (s *Option[T]) UnmarshalJSON(data []byte) error {

	if string(data) == "null" {
		s.isFull = false

		return nil
	}

	err := json.Unmarshal(data, &s.value)

	if err != nil {
		return err
	}

	s.isFull = true

	return nil
}

func Some[T any](value T) Option[T] {
	return Option[T]{value: value, isFull: true}
}

func None[T any]() Option[T] {
	return Option[T]{isFull: true}
}
