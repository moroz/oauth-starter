create extension if not exists citext with schema public;
create extension if not exists "uuid-ossp" with schema public;

CREATE OR REPLACE FUNCTION touch_updated_at() RETURNS trigger
   LANGUAGE plpgsql AS
$$BEGIN
   NEW.updated_at := now() at time zone 'utc';
   RETURN NEW;
END;$$;

create table users (
  id uuid primary key,
  email citext not null,
  inserted_at timestamp not null default (now() at time zone 'utc'),
  updated_at timestamp not null default (now() at time zone 'utc')
);

CREATE TRIGGER users_touch_updated_at BEFORE UPDATE ON users FOR EACH ROW EXECUTE PROCEDURE touch_updated_at();
