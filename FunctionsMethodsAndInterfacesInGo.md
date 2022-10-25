# Functions, Methods, and Interfaces in Go
## Course 2

### Module 1 Functions and organization
#### Objectives:
Identify the basic features and purposes of functions.
Identify the benefits of using pointers when working with functions.
Identify the difference between passing a slice and an array as function arguments.
Use functions and a slice to implement a sorting routine for integers.

assessments: 

*Notes:*
In a go program you must have a main function.
The main function is never called, it is called and invoked immediatly on run.
Functions allow reusability, abstraction.
A function can be static (have no parameter) or have parameter(s).

If you pass an array you have to copy it so it can take some time. Instead you can pass a pointer to an array. Actually you shall use slices instead.
Slices contain a pointer to an array so you can always use a slice instead of an array.
Passing a slice copies the pointer.

Guidelines for Functions:
- Function naming: pick a function name which describes behavior and which can be understood at once. Specify parameters names too.
- Functional Cohesion: a function shall perform only one operation. (ex: TriangleArea(), RectangleArea()). The code has to be "idiot-proof" (so simple anyone can understand and use it).
- Few parameters: easier debugging, more functional cohesion
- Function length: limit complexity. Use function Call Hierarchy (function calling a function etc.). Create smaller functions which are called by an other one.
- Control-flow complexity.

<br>
 ***

### Module 2 Function Types
#### Objectives:
- Identify advanced types, properties, and uses of functions.
- Identify the output that would result from running a given code block containing functions.
- Develop a routine containing functions in Go that solves a practical physics problem.


*Notes:*
Functions are first-class values, it means they can be treated like the other type (variables can be declared as "function" type)
They can be passed as argument to other functions, they can be returned from functions.
When you declare a var as a Func you declare without the "()".
use "... " in the argument list to say that it can takes a variable number of arguement of the specified type
You can defer a function execution with "defer". It defers until the surrounding function completes. Arguments are evaluated immediatly but the call is deferred.

<br>
 ***

### Module 3 Go’s object orientation
#### Objectives:
- Identify the basic properties and uses of objects and classes.
- Identify the differences between structs (in Go) and classes (in an object-oriented language).
- Identify the uses of methods and different data types.
- Develop an interactive Go routine that uses classes, objects, instances, structs, and methods to query the properties of existing instances.

*Notes:*
Classes = collection of data fields and functions that share a well-defined responsibility.
Object = instance of a class
Encapsulation = Data can be protected from the programmer, data can be accessed only using methods of the class.
No "class" keyword in GO.
Associating methods with data.
Methods has a receiver type that it is associated with. Use dot notation to use the method.
Usually use Struct as receiver as you can combine data types.

*Encapsulation*
```go
package data
type Point struct { 
	x float64
	y float64
} 
func (p *Point) InitMe(xn, yn float64) {
	p.x = xn
	p.y = yn
}
func (p *Point) Scale(v float64) {
	p.x = p.x * v
	p.y = p.y * v
}
func (p *Point) PrintMe(){
	fmt.Println(p.x, p.y)
}
```
and in main:
```go
package main
func main() {
	var p data.Point
	p.InitMe(3, 4)
	p.Scale(2)
	p.PrintMe()
}
```

Point receivers: 
Receiver type is implicitly passed as an argument to the methods. Argument passing in Go is "copy by value" so the method can not modify the data inside the receiver value.
To be able to change the actual value, we have to "call by reference" and put a pointer to the object in the receiver.
- Receiver is passed implicitly as an argument to the method
- Method cannot modify the data inside the receiver
- Receiver can be a pointer to a type
- Call by reference, pointer is passed to the method
```go
func (p *Point) OffsetX(v float64) {
	p.x = p.x + v
}
```
no need to dereference the pointer inside the method. Go knows waht we are trying to do and do it.
no need to reference either:
```go
func main() {
	p := Point{3, 4}
	p.OffsetX(5)
	fmt.Println(p.x)
}
```
-> go compiler recognize what we are trying to do

BEST PRACTICE:
All methods for a type have pointer receivers, or
All methods for a type have non-pointer receivers
Mixing pointer/non-pointer receivers for a type will get confusing
Pointer receiver allows modification

<br>
 ***

### Module 4 Interfaces for abstraction

#### Objectives:
- Identify the basic characteristics and features of interfaces in Go.
- Identify the characteristics of polymorphism.
- Identify the features associated with overriding and inheritance.
- Develop an interactive Go routine that uses classes, objects, instances, structs, and methods to create new class instances and later query their properties.

assessments: 

*Notes:*
Inheritance: subclass inherit the methods/data fo hte superclass.
Interfaces: concept useflul for polymorphism. An interface is a set of method signatures. Used to express conceptual similarity between types.
So if a type actually implements all the methods in the interface with the same method signatures, so same arguments, same name, same return values, then that type is said to satisfy the interface.
Defining an Interface type:
```go
type Shape2D interface {
   Area() float64
   Perimeter() float64
}
type Triangle {…}
func (t Triangle) Area() float64 {…}
func (t Triangle) Perimeter() float64 {…}
```
-> Triangle type satisfies the Shape2D interface (no need to state it explicitly)


<br>
 ***

**Glossary:**
function = set of instructions with a name
call by value = the arguments that are passed as parameters are copied to the parameters. It uses a *copy of the original*, not the original itself.
call by reference = instead of passing the argument, you pass a pointer (*int). Because you want the function to alter the elements passed to it.
closure = function + its environment
Classes = collection of data fields and functions that share a well-defined responsibility.
Object = instance of a class
Encapsulation = Data can be protected from the programmer, data can be accessed only using methods of the class.
Polymorphism = ability for an Object to have different "forms" depending on the context
Interfaces: set of method signatures.

------------
Template:
### Module 1 
#### Objectives:
assessments: 
*Notes:*
**Glossary:**