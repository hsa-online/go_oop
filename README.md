## Inheritance and "pure virtual" functions in Go ##

This article tries to answer the question: whether it is possible to simulate "the classic OOP" in Go. 

And the short answer is: *yes, it's possible*. But in result we get lots of clumsy boilerplate code, 
so the author warns that it’s more a theoretical possibility than a practically usable solution. 
Actually this code is just a first idea born during discussion of this article’s main question with one of my colleagues. 
Use it at your own risk :).

Assume we need to have this class hierarchy:

<p align="center">
  <img width="476" height="248" src="https://github.com/hsa-online/go_oop/blob/main/blob/oop_classes.png">
</p>

We know that there are no classes and no inheritance in Go, but internally the OOP concept is nothing more than a structure 
with fields and dispatch tables containing function pointers. From the other side in Go we have structs and first class 
functions (we can assign functions to variables). So in theory nothing prevents us from storing both the fields and dispatch 
table in the same struct. Also we will need to link structs to simulate inheritance, this adds one extra field into each “child” structure.

To keep the method presented more or less simple we allow to have only a single level of Parent->Child inheritance. 
Multiple levels of inheritance will require more complex logic to keep the correct instance pointer in the middle of hierarchy. 
Even with this simplification we still getting structures with lots of boilerplate:

<p align="center">
  <img width="696" height="304" src="https://github.com/hsa-online/go_oop/blob/main/blob/oop_structs.png">
</p>

In `Shape` we want to have a “pure virtual” function `GetArea() float64` which is later implemented in both `Circle` and `Square`. 
Also on any instance of the Shape we should be able to call `IsLargerThan(area float64) bool` which then calls the correct version of `GetArea()`.

### Additional note ###

From my point of view lots of boilerplate code is not as bad as the necessity of specifying an instance pointer everywhere 
(as a first parameter of each function call). Uniform Function Call Syntax (UFCS) may help here, but unfortunately 
we don’t have it in Go (sometimes UFCS helps to write elegant code 
[for example](https://en.wikipedia.org/wiki/Uniform_Function_Call_Syntax) in Nim).
