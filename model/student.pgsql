CREATE DATABASE my_db;

CREATE TABLE student (
    StdID int NOT NULL, 
    FirstName varchar(45) not null,
    LastName varchar(45) default null,
    Email varchar(45) not null,
    Primary key (StdID),
    Unique(Email)
)

truncate table student

DROP TABLE student
DROP TABLE admin
DROP TABLE enroll
DROP TABLE course

DROP DATABASE my_db