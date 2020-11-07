## Surin Lang

<img src="https://github.com/HSNHK/surin-lang/blob/master/resources/logo.png" width="250" >
<b>Surin programming language</b>

surin is an interpreter language written in golang or python
<br>
This means that its structure can be implemented in other languages.<br>
In practice, this language is just a structure and more work has been done on its structure
### instruction
```

code --> analyze code --> check pattern --> syntax function --> interpreter functions

```
In the code analysis section,<br>
if it matches the specified patterns,<br>
it will be executed<br>

Regexis used to analyze the code,
and each command has a pattern.<br>
For example, the print command has a pattern like this:
```

print("Hello world")
^[(print)|(\sprint)|(print\s)]+[(]+["].*["]+[)]$

```
## How to use it?

```

surin file.sur

or 

>>~#> code...

```
## syntax

```
[*]General functions

print("hello") #print string

print(1234)    #print int

logic(3,>=,2)  #logic functions (>,<,<=,>=,==)

math(2,**,4)   #math  functions (+,-,*,/,**)

streql("test","test") #Equality of two disciplines math(str1,str2)

len("hello programer") #string len function

find("hello","he") #find string to string

info 

time

[*]register

var(variableName) #create variables

push("variableName",value) #push to variable

pop(variableName) #get variable value

rm(variableName) #delete variable in stack

add(variableName,value) #add value to variable | +=

cmp(variable-1,variable-2)

type(variableName) #variable type

id(variableName) #variable id

[*]if

if(condition)?body:else

if(2,>=,2)?print("yes"):print("no")

[*]list

list(list-name:size) #create list

list(list-name,index,value) # set value for list with index

list(list-name,index) #return value with index

list(list-name) # print all list item [ index : value ]

//binary search
list.search(list-name,value) # search method for Returns its number if available otherwise -1

``` 
## MakeFile
```
make linux or windows

make help

make run or build

make compile-windows or compile-linux or compile-freebsd

make compile-all
```
## Documents

All documentation in the folder documentation

## Link
https://github.com/HSNHK/surin-lang
<br>
https://github.com/HSNHK/