## Surin Lang

<img src="https://github.com/HSNHK/surin-lang/blob/master/resources/logo.png" width="300" >
<br>
<b>Surin programming language</b>

surin is an interpreter language written in golang or python
<br>
This means that its structure can be implemented in other languages.

### instruction
```

code --> analyze code --> interpreter functions

```
In the code analysis section,<br>
if it matches the specified patterns,<br>
it will be executed<br>

Regexis used to analyze the code,
and each command has a pattern.
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
print("hello") #print string

print(1234)    #print int

logic(3,>=,2)  #logic functions (>,<,<=,>=,==)

math(2,**,4)   #math  functions (+,-,*,/,**)

streql("test","test") #Equality of two disciplines math(str1,str2)

len("hello programer") #string len function

``` 
## Documents

All documentation in the folder documentation