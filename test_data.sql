drop table if exists account;

create table account
(
    id       integer not null
        constraint accounts_pk
            primary key autoincrement,
    email    varchar not null,
    password varchar not null
);

create unique index accounts_email_uindex
    on account (email);

create unique index accounts_email_uindex_2
    on account (email);

create unique index accounts_id_uindex
    on account (id);

drop table if exists project;

create table project
(
    id      integer not null
        constraint project_pk
            primary key autoincrement,
    name    varchar not null,
    id_user integer
);

create unique index project_id_uindex
    on project (id);

create unique index project_id_user_name_uindex
    on project (id_user, name);

drop table if exists task;

create table task
(
    id         integer not null
        constraint task_pk
            primary key autoincrement,
    text       varchar not null,
    priority   integer not null,
    deadline   varchar,
    checked    boolean default false not null,
    id_project integer not null
);

create unique index task_id_project_priority_uindex
    on task (id_project, priority);

create unique index task_id_uindex
    on task (id);

INSERT INTO account (id, email, password) VALUES (1, 'user_for_tests@test.com', '$2a$10$.N0MOQ6ZP0pYGrk6QBZyWOuS5gZcfAN5kTFt0aEA4I8xMWYbo1ohi');
INSERT INTO project (id, name, id_user) VALUES (1, 'Project 1', 1);
INSERT INTO project (id, name, id_user) VALUES (2, 'Project 2', 1);
INSERT INTO task (id, text, priority, deadline, checked, id_project) VALUES (1, '1 Task for project 1', 1, '2020-10-31', 0, 1);
INSERT INTO task (id, text, priority, deadline, checked, id_project) VALUES (2, '2 Task for project 1', 2, '2020-10-31', 0, 1);
INSERT INTO task (id, text, priority, deadline, checked, id_project) VALUES (3, '3 Task for project 1', 3, '2020-10-31', 0, 1);
INSERT INTO task (id, text, priority, deadline, checked, id_project) VALUES (4, '4 Task for project 1', 4, '2020-10-31', 0, 1);
INSERT INTO task (id, text, priority, deadline, checked, id_project) VALUES (5, '1 Task for project 2', 1, '2020-10-31', 0, 2);
INSERT INTO task (id, text, priority, deadline, checked, id_project) VALUES (6, '2 Task for project 2', 2, '2020-10-31', 0, 2);
INSERT INTO task (id, text, priority, deadline, checked, id_project) VALUES (7, '3 Task for project 2', 3, '2020-10-31', 0, 2);
INSERT INTO task (id, text, priority, deadline, checked, id_project) VALUES (8, '4 Task for project 3', 4, '2020-10-31', 0, 2);