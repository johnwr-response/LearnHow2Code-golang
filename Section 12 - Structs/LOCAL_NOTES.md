[back](../LOCAL_NOTES.md)

## Section 12 - Structs
### Struct
```
md struct
cd struct
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/12/struct
go mod tidy
go run .
```
- A struct is a composite data structure (composite data types, aka, aggregate data types, aka, complex data types). Structs allows us to compose together values of different types.

### Embedded structs
```
md embedded-structs
cd embedded-structs
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/12/embedded-structs
go mod tidy
go run .
```
- A struct is a composite data structure (composite data types, aka, aggregate data types, aka, complex data types). Structs allows us to compose together values of different types.

### Reading documentation
### Anonymous structs
```
md anonymous-structs
cd anonymous-structs
ni main.go
Add-Content .\main.go "package main`n`nimport `"fmt`"`n`nfunc main() {`n`tfmt.Println(`"Hello code`")`n}"
go mod init golang-course/12/anonymous-structs
go mod tidy
go run .
```
### Housekeeping
  - It's ALL about type
    - Is go an object orientated language?
      - Yes and no.
      - Go has `OOP` aspects, but it's all about `TYPE`.
      - We create `TYPES` in go; user-defined `TYPES`.
      - We can then have `VALUES` of that `TYPE`. We can assign `VALUES` of a user-defined `TYPE` to `VARIABLES`.
  - Go IS Object Oriented
    1. Encapsulation
      - a) state (`fields`)
      - b) behaviour (`methods`)
      - c) exported/unexported; viewable/non-viewable
    1. Reusability
      - a) inheritance (`embedded types`)
    1. Polymorphism
      - a) interfaces
    1. Overriding
      - a) `promotion`
  - Traditional OOP
    1. Classes
      - a) data structure describing a type of object
      - b) you can then create `instances`/`objects` from the class/blueprint
      - c) classes hold both:
        1. state/data/fields
        1. behavior/methods
      - d) `public`/`private`
    1. Inheritance
  - In GO
    1. You don't create classes, you create a `TYPE`
    1. You don't instantiate, you create a `VALUE` of a `TYPE`
  - User defined types
    - We can declare a new type `foo`
      - the `underlying type` of foo is an `int`
      - int conversion
        - int(myAge)
        - converting type `foo` to type `int`
    - It is a ***BAD PRACTICE*** to ***ALIAS TYPES***
      - One exception
        - If you need to attach methods to a type
        - See the time package for an example of this [godoc.org/time](https://pkg.go.dev/time)
          - type Duration int64
          - Duration has methods attached to it
  - Named types vs Anonymous types
    - Anonymous types are indeterminate. They have not been declared as a type yet. The compiler has flexibility with anonymous types. You can assign an anonymous type to a variable declared as a certain type. If the assignment can occur, the compiler will figure it out; the compiler will do an implicit conversion. You cannot assign a named type to a different named type.
  - Padding & archtectural alignment
    - Convention: logically organize your fields together. Readability & clarity trump performance as a design concern. Go will be performant. Go for readability first. However, if you are in a situation where you need to prioritize performance: lay the fields out from largest to smallest, eg, int 64, int64, float32, bool  