create table enroll (
    std_id int not null,
    course_id int not null,
    date_enrolled varchar(45) default null,
    primary key (std_id,course_id),
    constraint course_fk foreign key (course_id) references course 
    (CourseId) on delete cascade on update cascade,
    CONSTRAINT std_fk foreign key (std_id) references student
    (StdID) on delete cascade on update cascade
)