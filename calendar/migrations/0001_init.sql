-- CREATE USER event_admin
-- LOGIN
-- password '123';
--
-- create database eventsdb owner event_admin;

-- auto-generated definition
create table events
(
    id          bigserial             not null
        constraint events_pk
            primary key,
    name        varchar(256)          not null,
    description varchar,
    date        timestamp             not null,
    duration    bigint                not null,
    deleted     boolean default false not null
);

alter table events
    owner to event_admin;

create unique index events_id_uindex
    on events (id);


