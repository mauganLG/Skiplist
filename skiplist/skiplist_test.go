package skiplist

import (
	"reflect"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkipStr(t *testing.T) {
	list := NewSkipListD[string]()

	list.Insert(3, "3")
	list.Insert(6, "6")
	list.Insert(9, "9")

	value, _ := list.Search(6)

	assert.Equal(t, "6", value)
}

func TestSkipInt(t *testing.T) {
	list := NewSkipListD[int]()

	list.Insert(1, 23)
	list.Insert(2, 78)
	list.Insert(3, 2)

	value, _ := list.Search(2)

	assert.Equal(t, 78, value)

}

func TestSkipUpdate(t *testing.T) {
	list := NewSkipListD[int]()

	list.Insert(12, 23)
	list.Insert(1, 78)
	list.Insert(5, 2)

	value, _ := list.Search(5)

	assert.Equal(t, 2, value)

	list.Insert(5, 123)

	value, _ = list.Search(5)

	assert.Equal(t, 123, value)

}

func TestSkipSearchFalse(t *testing.T) {
	list := NewSkipListD[rune]()

	list.Insert(10, 'a')
	list.Insert(14, 'P')
	list.Insert(23, '&')

	_, res := list.Search(22)

	assert.Equal(t, false, res)

}

func TestSkipDelete(t *testing.T) {
	list := NewSkipListD[float32]()

	list.Insert(4, 2.3)
	list.Insert(9, 7.8)
	list.Insert(12, 0.9)

	assert.Equal(t, true, list.Delete(4))

}

func TestSkipDeleteFalse(t *testing.T) {
	list := NewSkipListD[float64]()

	list.Insert(4, 2.3)
	list.Insert(9, 7.8)
	list.Insert(12, 0.9)

	assert.Equal(t, false, list.Delete(6))

}

func TestSkipDeleteFalseEmpty(t *testing.T) {
	list := NewSkipListD[float32]()

	assert.Equal(t, false, list.Delete(0))

}

func TestSkipPerson(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	slPeople := NewSkipListD[Person]()
	slPeople.Insert(1, Person{Name: "Alice", Age: 30})
	slPeople.Insert(2, Person{Name: "Bob", Age: 25})

	slicesPerson := slices.Collect[Person](slPeople.Values())

	P := []Person{
		{
			Name: "Alice",
			Age:  30,
		},
		{
			Name: "Bob",
			Age:  25,
		},
	}
	res := reflect.DeepEqual(slicesPerson, P)

	assert.EqualValues(t, true, res)

}

func TestSkipLength(t *testing.T) {
	type Car struct {
		Model string
		Year  int
	}

	slCar := NewSkipListD[Car]()
	slCar.Insert(1, Car{Model: "205", Year: 1982})
	slCar.Insert(2, Car{Model: "Countach", Year: 1974})
	slCar.Insert(3, Car{Model: "Testarossa", Year: 1984})

	assert.EqualValues(t, uint(3), slCar.Length())

	slCar.Delete(2)
	assert.EqualValues(t, uint(2), slCar.Length())
}
