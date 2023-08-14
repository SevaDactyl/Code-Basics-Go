/*
Реализуйте интерфейс IVoiceable для структур Cat, Cow и Dog так, чтобы при вызове метода Voice экземпляр структуры Cat возвращал строку "Мяу", экземпляр Cow строку "Мууу", а экземпляр Dog сообщение Гав:

cat := Cat{} 
dog := Dog{}
cow := Cow{}

fmt.Println(cat.Voice()) // Мяу
fmt.Println(dog.Voice()) // Гав
fmt.Println(cow.Voice()) // Мууу
*/

package solution

type IVoiceable interface {
    Voice() string
}

type Cat struct {
    // … 
}

type Cow struct {
    // … 
}

type Dog struct {
	// … 
}

// BEGIN (write your solution here)

func (c Cat) Voice() string {
	return "Мяу"
}

func (cw Cow) Voice() string {
	return "Мууу"
}

func (d Dog) Voice() string {
    return "Гав"
}

// END
