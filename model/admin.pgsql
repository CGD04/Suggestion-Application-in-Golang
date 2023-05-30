create table admin (
    FirstName varchar(45) not null,
    LastName  varchar(45) default null,
    Email varchar(45) not null,
    Password varchar(45) not null,
    Primary key (Email)
)