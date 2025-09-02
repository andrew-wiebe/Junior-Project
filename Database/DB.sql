
-- Note: uuids are intentionally varchars, this lets us expose them to the user
Create table users(
    user_uuid varchar(255) primary key not null,
    username varchar(255) not null,
    psswrd varchar(255) not null,
    salt varchar(255) not null, -- salt for passwords
)

Create table course(
    couse_uuit varchar(255) primary key not null,
    code varchar(255) not null,
    title varchar(255) not null,
)

Create table professor(
    prof_uuid varchar(255) primary key not null,
    prof_name varchar(255) not null,
)

Create table course_rating(
    course_rating_uuid varchar(255) primary key not null,
    poster_uuid varchar(255) not null references users.user_uuid,
    stars int not null, -- This and hours/week could be nullable
    hours_per_week float not null,
    teacher_uuid references professor.prof_uuid,
    review_text varchar(400),
)