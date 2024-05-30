SOLID Design Principles:

The solid design principle is a set of 5 guidelines that help software developers design system that are easy to maintain , extend and understand. Introduced by Robert C. Martin(Uncle Bob)  are widely used and regarded as best practise in object-oriented software development. 

1. Single Responsibility Principle : A class should have only one reason to change , meaning it should have one responsibility or job . 
    E.g. : A class responsible for managing journals should not handle I/O or persistence operations . 
    Idea is to do composition so that code can be clear separation , reused , can be tested independently.

2. open - closed principle : It states that software entities should be open to extension but should be closed to modification. 
    The design and writing of the codes should be done in such a way that new functionality should be added with minimum number of changes in the existing code.
    The design should be done in a way to allow the adding of new functionality as new classes , keeping as much as possible existing code unchanged.

    Explanation : 
    -> Open for extension : You can add new functionality to a module
    -> Closed for modification: You should not change existing code when you add new functionality.

In practical terms :
-> When you need to add new features , you should not modify existing struct/classes. Instead you should create new struct/classes or interfaces that extend existing functionality.
->This ensures that existing code remains stable and unaltered, reducing the risk of introducing bugs.


3. Liskov Substitution Principle

A corrct way to do inheritance 

Client should not know which specific subtype they are calling
->All the base class method must be applicable to derived class methods.

->No new exception should be thrown by subclasses.
Iemployee interface has GetMinimumSalary()
IEmployeeBonus interface has CalBonus
Employee class implments IEmployee and IEmployeeBonus

PermanentEmployee class extends Employee
Templorary class inherits Employee
ContractEMp Class only inherits IEmployee since COntract employee donot have bonus.


4. Interface Segregation Principle.

"Many client specific interfaces are better than general-purpose interface"
We should not enforce clients to implement interfaces that they don't use. Instead of creating one interface we can break down it to smaller interfaces.

5. Dependency Inversion Principle.
High-level modules should not depend on low level modules. Both should depend on abstractions
Abstraction should not depend on details. Details should be dependent on abstraction.
The interaction between high level and low level modules should be thought of as an abstract interaction between them.

Basically , DIP states that Low-level class object should not be directly created in main class , make it abstract relation between and then call it in as your local thing.


----------------------------------------------------------------------------------------------------------------------------------------------------------------------


If we don't follow SOLID Principle:

End up with tight or strong coupling of the code with many other module/applications

Tight coupling causes time to implement any new requirement, features or any bug fixes and sometimes it create unknown issues

End up with a code that is not testable

End up with duplication of code

End up creating new bugs by fixing another bug

End up with many unknown issues in the application development life cycle.


Benefits :

Achieve reduction , complexity of code
Increase readability  extensibility and maintenance
Reduce errors and implement reusability
Achieve better testability
Reduce tight coupling

Solution for successful application :
1.Architecture e.g. MVC , MVVM 
2.Design Principles
3.Design Patterns

     
    
    
    
    
    
