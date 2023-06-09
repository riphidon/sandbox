-- Database: sandbox

-- DROP DATABASE IF EXISTS sandbox;

CREATE DATABASE sandbox
    WITH 
    OWNER = user
    ENCODING = 'UTF8'
    LC_COLLATE = 'en_US.UTF-8'
    LC_CTYPE = 'en_US.UTF-8'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

-- Table: public.items

-- DROP TABLE IF EXISTS public.items;

CREATE TABLE IF NOT EXISTS public.items
(
    item_id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    item_name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    description character varying(255) COLLATE pg_catalog."default",
    maker_name character varying(100) COLLATE pg_catalog."default" NOT NULL,
    category character varying(50) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT items_pkey PRIMARY KEY (item_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.items
    OWNER to user;


-- Table: public.makers

-- DROP TABLE IF EXISTS public.makers;

CREATE TABLE IF NOT EXISTS public.makers
(
    id_maker integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    maker_name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT makers_pkey PRIMARY KEY (id_maker)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.makers
    OWNER to user;

-- Table: public.users

-- DROP TABLE IF EXISTS public.users;

CREATE TABLE IF NOT EXISTS public.users
(
    user_id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    first_name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    last_name character varying(50) COLLATE pg_catalog."default" NOT NULL,
    nickname character varying(25) COLLATE pg_catalog."default" NOT NULL,
    email character varying(50) COLLATE pg_catalog."default" NOT NULL,
    password_hash character varying(255) COLLATE pg_catalog."default" NOT NULL,
    remember character varying(255) COLLATE pg_catalog."default" NOT NULL,
    remember_hash character varying(255) COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp with time zone NOT NULL DEFAULT CURRENT_TIMESTAMP,
    CONSTRAINT users_pkey PRIMARY KEY (user_id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.users
    OWNER to user;

-- Table: public.users

-- TRIGGER to prevent creation timestamp overwrite

CREATE OR REPLACE FUNCTION stop_change_on_timestamp()
  RETURNS trigger AS
$BODY$
BEGIN
  -- always reset the timestamp to the old value ("actual creation time")
  NEW.timestamp := OLD.timestamp;
  RETURN NEW;
END;
$BODY$
language 'plpgsql';


CREATE TRIGGER prevent_timestamp_changes
  BEFORE UPDATE
  ON users
  FOR EACH ROW
  EXECUTE PROCEDURE stop_change_on_timestamp();
