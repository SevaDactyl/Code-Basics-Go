/*
Представим, что есть структура Person, содержащая возраст человека:

type Person struct {
    Age uint8
}
Реализуйте тип PersonList (слайс структур Person), с методом (pl PersonList) GetAgePopularity() map[uint8]int, который возвращает мапу, где ключ — возраст, а значение — кол-во таких возрастов:

pl := PersonList{
  {Age: 18},
  {Age: 44},
  {Age: 18},
}

pl.GetAgePopularity() // map[18:2 44:1]
*/

package solution

// Person is a struct that keeps info about person's age
type Person struct {
	Age uint8
}

// BEGIN (write your solution here)

// PersonList is a list of persons.
type PersonList []Person

// GetAgePopularity calculates and returns popularity for each age in the person list.
func (pl PersonList) GetAgePopularity() map[uint8]int {
	popularity := make(map[uint8]int)
	for _, p := range pl {
		popularity[p.Age]++
	}

	return popularity
}

// END
