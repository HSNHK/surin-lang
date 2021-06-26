## Surin Lang [![HSNHK - surin-lang](https://img.shields.io/static/v1?label=HSNHK&message=surin-lang&color=blue&logo=github)](https://github.com/HSNHK/surin-lang)
[![stars - surin-lang](https://img.shields.io/github/stars/HSNHK/surin-lang?style=social)](https://github.com/HSNHK/surin-lang)
[![forks - surin-lang](https://img.shields.io/github/forks/HSNHK/surin-lang?style=social)](https://github.com/HSNHK/surin-lang)
[![GitHub release](https://img.shields.io/github/release/HSNHK/surin-lang?include_prereleases=&sort=semver)](https://github.com/HSNHK/surin-lang/releases/)
[![License](https://img.shields.io/badge/License-MIT-blue)](#license)
[![issues - surin-lang](https://img.shields.io/github/issues/HSNHK/surin-lang)](https://github.com/HSNHK/surin-lang/issues)

<img src="https://github.com/HSNHK/surin-lang/blob/master/resources/Capture.JPG">
<b><i>Surin programming language</i></b>

surin is an interpreter language written in golang.
<br>
This commentator has a different structure than other commentators and is `designed for testing only`.

### Functional form
```

[Code] --> Code analysis --> Find the pattern --> Delivery to function --> Execute the command

```

## How to use it ?

```

surin file.sur

* or Enter active mode *

>>~#> code...

```
## syntax
| Name | Description | Sample |
|------|-------------|--------|
| `Print` | This function is used to print `strings` and its writing form is similar to the `print function in Python` | ```print("hello")``` |
| `Print` | This function is used to print `integer` and its writing form is similar to the `print function in Python` | ```print("hello")``` |
| `Logic` | This function is used for `logical` operators. The first number is in the first argument and the logical operator is in the second argument, and we enter the second number in the third argument. Authorized operators (`>`,`<`,`<=`,`>=`,`==`) | ```logic(3,>=,2)``` |
| `Math` | This function is used to perform `mathematical` calculations.  Authorized operators (`+`,`-`,`*`,`/`,`**`) | ```math(2,**,4)``` |
| `Streql` | This function is used to check the `equality of two strings`. | ```streql("strA","strB")``` |
| `Len` | This function is used to calculate the length of a string. | ```len("hello programer")``` |
| `Find` | This function is used to find a string in another string. | ```find("hello","he")``` |
| `Info` | This operator is for obtaining interpreter and system specifications. | ```info``` |
| `Time` | This operator is used to obtain the current system time. | ```time``` |
| `Var` | This function is used to construct a variable. | ```var(variableName)``` |
| `Push` | This function is used to set a variable. A variable can store any kind of value. | ```push("variableName",value)``` |
| `Pop` | This function is used to get the last value in a variable. | ```pop(variableName)``` |
| `rm` | This function is used to delete a variable from memory. | ```rm(variableName)``` |
| `Add` | This function is used to add a value to a variable. | ```add(variableName,value)``` |
| `Cmp` | This function is used to check that the values ​​of two variables are equal. | ```cmp(variable-1,variable-2)``` |
| `Type` | This function is used to `display the value of a variable`. | ```type(variableName)``` |
| `Id` | This function is used to get the `address of a variable`. | ```id(variableName)``` |
| `If` | This command is used to create a `condition`. | ```if(2,>=,2)?print("yes"):print("no")``` | 
| `List` | This function is used to create a list. The name and length of the list are given to the first argument | ```list(list-name:size)``` |
| `List` | To set an value of a list, enter the name of the list in the first argument, the index of the section in the second argument, and the desired value in the third argument. | ```list(list-name,index,value)``` |
| `List` | To get the value of a certain part of the list, you must give the name and number of the desired part to this function. | ```list(list-name,index)``` |
| `List` | To display all the values ​​in the list, just name it to this function. | ```list(list-name)``` |
| `list.search` | You can use this function to search the list. `binary search` | ```list.search(list-name,value)``` |
| `list.sort` | You can use this function to sort items in a list. `bubble sort` | ```list.sort(list-name)``` |
| `list.del` | You can use this function to delete a list from memory. | ```list.del(list-name)``` |
## MakeFile

`make file for linux and windows`
```
make help

make run or build

make compile-windows
make compile-linux
make compile-freebsd

make compile-all
```
## License

Released under [MIT](/LICENSE) by [@HSNHK](https://github.com/HSNHK).
**This project is no longer being developed**
