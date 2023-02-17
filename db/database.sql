create database api_todo;

create type state as enum ('active', 'completed');

create table todo (
    id serial primary key,
    description text,
    title text,
    state state,
    created_at timestamp default now(),
    updated_at timestamp default now()
);