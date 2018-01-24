# API Implementation of Distant School
## Scheme

A list of entities who are first-citizen users:
* Student
* Teacher

In order to put a mark, it was decided to divide mark and subject types:
* Mark
* Subject
* _Date_

Date type calls _Mark_ and _Subject_ types and bind them into an __Event__. Event is just a nice property for an individual mark which makes it by design flexible and clear for both _student_ and _teacher_.
