CREATE TABLE IF NOT EXISTS migrations
(
    id primary key not null,
    name varchar(255) not null,
    time timestamp not null
);

CREATE TABLE IF NOT EXISTS scheduled_notes
(
    id integer not null,
    chat_id integer not null,
    message_id integer not null,
    message text not null,
    repeats integer not null,
    type varchar(255) null
);
