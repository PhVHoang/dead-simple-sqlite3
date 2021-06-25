# dead-simple-sqlite3 (this side project is undergoing...)


## Sqlite?
* A database engine
* Allows users to interact with a relational database
* **A database is stored in a single file**

## Sqlite structure
![Screen Shot 2021-06-23 at 17 27 39](https://user-images.githubusercontent.com/12546802/123063403-555faf00-d448-11eb-9093-6a2f90df76cf.png)

## REPL
- [x] Added simple REPL
<img width="658" alt="Screen Shot 2021-06-23 at 23 56 07" src="https://user-images.githubusercontent.com/12546802/123120883-a5a63380-d47f-11eb-83c8-796578945df6.png">

## Barebone Core/Compiler
<img width="648" alt="Screen Shot 2021-06-25 at 0 24 10" src="https://user-images.githubusercontent.com/12546802/123290024-e7041500-d54b-11eb-872f-f7054a8e1645.png">

## Replace naive Compiler with sqlparser
Replace naive implemeted compiler with this [sqlparser](https://github.com/marianogappa/sqlparser) library.

<img width="653" alt="Screen Shot 2021-06-26 at 0 39 52" src="https://user-images.githubusercontent.com/12546802/123450216-76c2c580-d617-11eb-98d7-e88b4d507f5f.png">

## Implement `insert`
<img width="652" alt="Screen Shot 2021-06-26 at 1 46 27" src="https://user-images.githubusercontent.com/12546802/123458862-7d096f80-d620-11eb-9739-88cd6e57fc12.png">


## TODO
- [x] Core structure
- [x] SQL Parser
- [x] Basic statement (insert)
- [] ...
## References
This project is heavily borrowed from
1. https://cstack.github.io/db_tutorial/parts/part1.html
2. https://www.youtube.com/watch?v=yFGPiftpIJY
