# Design.Patterns.Examples
Implementing design patterns for fun and learning


# SOLID<br>
### Single Responsibility Principle 
Keep object as simple as possible with segregated responsibilities that belongs to this objec. <BR> Like for example Blog struct is responsible for menaging blog posts with Add, Update, Remove.. functions so don't implement anti pattern like God Object that is doing everything... menaging blog posts, writing, reading from a file etc.  This persistence functionality can be separated with new struct or library usage. 