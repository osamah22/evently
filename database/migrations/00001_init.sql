-- +goose Up
CREATE TABLE if not exists users (
    id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
    auth0_id TEXT NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW()
);

create table if not exists categories (
    id serial primary key,
    name varchar(100) not null unique,
    created_at timestamptz not null default now()
);

insert into categories (name) values
    ('MUSIC'), ('SPORTS'), ('TECH'), ('ART'), ('FOOD'),
    ('BUSINESS'), ('EDUCATION'), ('HEALTH'), ('NETWORKING'),
    ('COMMUNITY'), ('OTHER');

create table if not exists events (
    id uuid primary key,
    title varchar(512) not null,
    description varchar(1024) not null,
    status varchar(256) not null check (status in ('PUBLISHED', 'DRAFT', 'CANCELLED')),
    category_id int not null references categories(id),
    organizer_id uuid not null,
    location varchar(255),
    starting_time timestamptz not null,
    ending_time timestamptz not null,
    check(starting_time < ending_time),
    created_at timestamptz not null default now(),
    updated_at timestamptz not null default now()
);

create index if not exists events_status_idx on events(status);
create index if not exists events_category_id_idx on events(category_id);

-- +goose Down
drop table if exists events;
drop table if exists categories;
drop table if exists users;
