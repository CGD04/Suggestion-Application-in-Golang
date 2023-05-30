
create table course (
    CourseId int not null,
    CourseName varchar(45) not null,
    Primary key (CourseId)
)
alter table course
alter column CourseId type varchar(45);