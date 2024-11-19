package skiplist

import (
	"reflect"
	"slices"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSkipStr(t *testing.T) {
	list := NewSkipList[string]()

	list.Insert(3, "3")
	list.Insert(6, "6")
	list.Insert(9, "9")

	value, _ := list.Search(6)

	assert.Equal(t, "6", value)
}

func TestSkipInt(t *testing.T) {
	list := NewSkipList[int]()

	list.Insert(1, 23)
	list.Insert(2, 78)
	list.Insert(3, 2)

	value, _ := list.Search(2)

	assert.Equal(t, 78, value)

}

func TestSkipPerson(t *testing.T) {
	type Person struct {
		Name string
		Age  int
	}

	slPeople := NewSkipList[Person]()
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

	slCar := NewSkipList[Car]()
	slCar.Insert(1, Car{Model: "205", Year: 1982})
	slCar.Insert(2, Car{Model: "Countach", Year: 1974})
	slCar.Insert(3, Car{Model: "Testarossa", Year: 1984})

	assert.EqualValues(t, slCar.len(), slCar.Length())

	slCar.Delete(1)
	assert.EqualValues(t, slCar.len(), slCar.Length())
}
