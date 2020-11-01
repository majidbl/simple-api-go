# simple-api-go
simple university api in golang

# TODO
- [x] add rest api for teachers
- [x] add rest api for students
- [x] add rest api for teacher
- [x] add simple unit testing
- [ ] add jwt auth
- [ ] use redis for data extract
## Usage
###### GET     http://localhost:4321/api/v1.0/students/      get list students
###### GET     http://localhost:4321/api/v1.0/students/:sid  get specific students
###### POST    http://localhost:4321/api/v1.0/students/      create new students
###### UPDATE  http://localhost:4321/api/v1.0/students/:sid  update specific students
###### DELETE  http://localhost:4321/api/v1.0/students/:sid  delete specific students
###### GET     http://localhost:4321/api/v1.0/teachrs/       get list teacher
###### GET     http://localhost:4321/api/v1.0/teacher/:sid   get specific teacher
###### POST    http://localhost:4321/api/v1.0/teacher/       create new teacher
###### UPDATE  http://localhost:4321/api/v1.0/teachers/:sid  update specific teachers
###### DELETE  http://localhost:4321/api/v1.0/teachers/:sid  delete specific teachers

###### GET     http://localhost:4321/api/v1.0/class/students/:sid       get all class for specific students
###### GET     http://localhost:4321/api/v1.0/class/students/:sid/:cid  get specific class for specific students
###### POST    http://localhost:4321/api/v1.0/class/students/:sid/:cid  add new class for students
###### DELETE  http://localhost:4321/api/v1.0/class/students/:sid/:cid  delete specific class for specific students


###### GET     http://localhost:4321/api/v1.0/class/teachrs/:tid            get all class that created by specific teacher
###### GET     http://localhost:4321/api/v1.0/class/teacher/:tid/:cid       get specific class that created by specific teacher
###### POST    http://localhost:4321/api/v1.0/class/teacher/:tid            create new class by specific teacher
###### UPDATE  http://localhost:4321/api/v1.0/class/teachers/:tid           update specific class by teachers that created it
###### DELETE  http://localhost:4321/api/v1.0/class/teachers/:tid           delete specific class by teachers that created it
