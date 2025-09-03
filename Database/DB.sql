Drop table if exists course_shedule_association; 
Drop table if exists user_schedule;
Drop table if exists course_rating;
Drop table if exists professor;
Drop table if exists course;
Drop table if exists user;
-- Drop table if exists 
-- Drop table if exists 




-- Note: uuids are intentionally varchars, this lets us expose them to the user
Create table user(
    user_uuid varchar(255) primary key not null,
    username varchar(255) not null,
    psswrd varchar(255) not null,
    salt varchar(255) not null, -- salt for passwords
);

Create table course(
    couse_uuid varchar(255) primary key not null,
    code varchar(255) not null,
    title varchar(255) not null,
);

Create table professor(
    prof_uuid varchar(255) primary key not null,
    prof_name varchar(255) not null,
);


Create table course_rating(
    course_rating_uuid varchar(255) primary key not null,
    poster_uuid varchar(255) not null references user.user_uuid,
    course_uuid varchar(255) not null references course.course_uuid,
    stars int not null, -- This and hours/week could be nullable
    hours_per_week float not null,
    teacher_uuid references professor.prof_uuid,
    review_text varchar(400),
);

Create table user_schedule(
    user_schedule_uuid varchar(255) primary key not null,
    user_uuid varchar(255) not null references user.user_uuid,
    schedule_title varchar(200),
);

Create table course_shedule_association(
    course_shedule_association_uuid varchar(255) primary key not null,
    shedule_uuid varchar(255) not null references user_schedule.user_schedule_uuid,
    course_uuid varchar(255) not null references course.course_uuid
);