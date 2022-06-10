# Design.Patterns.Examples
Implementing design patterns for fun and learning


# SOLID<br>

### Single Responsibility Principle 
Keep object as simple as possible with segregated responsibilities that belongs to this objec. <BR> Like for example Blog struct is responsible for menaging blog posts with Add, Update, Remove.. functions so don't implement anti pattern like God Object that is doing everything... menaging blog posts, writing, reading from a file etc.  This persistence functionality can be separated with new struct or library usage. 


### Open Closed Principle 
Open for extension but closed for modyfication. Instead of modification of existing and tested code just add new struct with expected functionality that is extending functionality of existing code. The Specyfication pattern can help with that. 


### Liskov Substytution Principle
So what the list of substitution principle basically states that if you have some API that takes a base class and works correctly with that base class it should also work correctly with the derived class. Unfortunately in go we don't have base classes in derived classes. But let's focuse on what we can actualy implement in go!
Liskov principal basically states that if you continue to use generalizations like interfaces for example then you should not have inheritance or you should not have implementers of those generalizations break some of the assumptions which are set up at the higher level so if you have interface A and you swith to interface B the behaviour should stay the same in Go. <br> 
(It states that a superclass object should be replaceable with a subclass object without breaking the functionality of the software. It is a type of behavioral subtyping defined by semantic, rather than syntactic, design consideration.)

### Interface Segregation Principle
Easiest of the principles. Don't put everything in one interface - segregate responsibilities. 

### Dependency Inversion Principle

It's something different than Dependency Injection!!! They do have a relationship tho!
- High lvl modules shouldn't depend on low level modules.
- Both of them should depend on abstraction.

---
# Builder 
Builder pattern is great for building complicated objects. You simply building object piiece by piece adding neccecary functionality to it. 
Builder pattern is optimal solution when objects can't be created with one simple constructor call, requires a lot of ceremony over creating or you have to use factory with multiple arguments to create object. Builder pattern will should provide elegant way for constructing object step by step. 
- "When  piecewise object consturcion is complicated, provide an API for doing it succinctly" 
https://stackoverflow.com/questions/328496/when-would-you-use-the-builder-pattern

# Factory

Factory is a component responsible for wholesale creation of object (it's opposite to Builder since builder is constructing multiparameter object step by step).
- You know you need factory when creation logic becomes too complex (like for example struct witch has a lot of lists and properties in it). 
- In Go we use a lot of factories as constructors like in other programming languages.

# Adapter

Adapter is a structural design pattern, which allows incompatible objects to collaborate. The Adapter acts as a wrapper between two objects. It catches calls for one object and transforms them to format and interface recognizable by the second object.
