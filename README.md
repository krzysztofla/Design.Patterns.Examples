# Design.Patterns.Examples
Implementing design patterns for fun and learning


# SOLID<br>

### Single Responsibility Principle 
Keep object as simple as possible with segregated responsibilities that belongs to this objec. <BR> Like for example Blog struct is responsible for menaging blog posts with Add, Update, Remove.. functions so don't implement anti pattern like God Object that is doing everything... menaging blog posts, writing, reading from a file etc.  This persistence functionality can be separated with new struct or library usage. 


### Open Closed Principle 
Open for extension but closed for modyfication. Instead of modification of existing and tested code just add new struct with expected functionality that is extending functionality of existing code. The Specyfication pattern can help with that. 


### Liskov Substytution Principle
So what the list of substitution principle basically states that if you have some API that takes a base class and works correctly with that base class it should also work correctly with the derived class. Unfortunately in go we don't have base classes in derived classes. But let's focuse on what we can actualy implement in go!