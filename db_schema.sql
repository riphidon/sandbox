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
