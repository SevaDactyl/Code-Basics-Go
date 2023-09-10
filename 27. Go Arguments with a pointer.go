/*
Реализуйте функцию CopyParent(p *Parent) Parent, которая создает копию структуры Parent и возвращает ее:

type Parent struct {
    Name     string
    Children []Child
}

type Child struct {
    Name string
    Age  int
}

cp := CopyParent(nil) // Parent{}

p := &Parent{
   Name: "Harry",
   Children: []Child{
       {
           Name: "Andy",
           Age:  18,
       },
   },
}
cp := CopyParent(p)

// при мутациях в копии "cp"
// изначальная структура "p" не изменяется
cp.Children[0] = Child{}
fmt.Println(p.Children) // [{Andy 18}]
*/

package solution

// Parent is a parent struct.
type Parent struct {
	Name     string
	Children []Child
}

// Child is a child struct.
type Child struct {
	Name string
	Age  int
}

// BEGIN (write your solution here)

// CopyParent makes a copy of the Parent struct and returns it.
func CopyParent(p *Parent) Parent {
	if p == nil {
		return Parent{}
	}

	cp := *p

	cp.Children = make([]Child, len(p.Children))
	copy(cp.Children, p.Children)

	return cp
}

// END
