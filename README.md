# API Implementation of Distant School
## Scheme

A list of entities which are first-citizen users:
* Student
* Teacher

In order to put a mark for a teacher, it was decided to divide mark and subject types:
* Mark
* Subject
* _Date_

Date type calls _Mark_ and _Subject_ types and binds them into an __Event__. Event is just a nice property for an individual mark which makes it by design flexible and clear for importing to other databases or blockchains.
