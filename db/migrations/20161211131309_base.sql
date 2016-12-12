
-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied

--
-- PostgreSQL database dump
--

--SET statement_timeout = 0;
--SET client_encoding = 'UTF8';
--SET standard_conforming_strings = on;
--SET check_function_bodies = false;
--SET client_min_messages = warning;

--
-- Name: plpgsql; Type: EXTENSION; Schema: -; Owner:
--

--CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;


--
-- Name: EXTENSION plpgsql; Type: COMMENT; Schema: -; Owner:
--

--COMMENT ON EXTENSION plpgsql IS 'PL/pgSQL procedural language';
--
--
--SET search_path = public, pg_catalog;
--
--SET default_tablespace = '';
--
--SET default_with_oids = false;

--
-- Table Users
--
CREATE TABLE users (
    id integer NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    name character varying(255) NOT NULL,
    username character varying(30) NOT NULL,
    email character varying(255) NOT NULL,
    hashed_password bytea NOT NULL,
    password character varying(255),
    image character varying(255),
    curator boolean,
    admin boolean,
    email_verified_at timestamp with time zone
);

CREATE SEQUENCE users_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE users_id_seq OWNED BY users.id;

ALTER TABLE ONLY users ALTER COLUMN id SET DEFAULT nextval('users_id_seq'::regclass);

SELECT pg_catalog.setval('users_id_seq', 1, false);

ALTER TABLE ONLY users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);

CREATE UNIQUE INDEX uix_users_email ON users USING btree (email);
CREATE UNIQUE INDEX uix_users_username ON users USING btree (username);

-- --
-- -- Table Arguments
-- -- Stop here
-- CREATE TABLE arguments (
--     id integer NOT NULL,
--     created_at timestamp with time zone DEFAULT now(),
--     updated_at timestamp with time zone DEFAULT now(),
--     deleted_at timestamp with time zone,
--     title character varying(255) NOT NULL,
--     descriptrion character varying(4000),
--     truth numeric,
--     created_by integer,
--     argument_pro_id integer,
--     reference_id integer
-- );

-- CREATE SEQUENCE arguments_id_seq
--     START WITH 1
--     INCREMENT BY 1
--     NO MINVALUE
--     NO MAXVALUE
--     CACHE 1;

-- ALTER SEQUENCE arguments_id_seq OWNED BY debates.id;

-- ALTER TABLE ONLY arguments ALTER COLUMN id SET DEFAULT nextval('arguments_id_seq'::regclass);

-- SELECT pg_catalog.setval('arguments_id_seq', 1, false);

-- ALTER TABLE ONLY arguments
--     ADD CONSTRAINT arguments_pkey PRIMARY KEY (id);

--
-- Table Debates
--
CREATE TABLE debates (
    id integer NOT NULL,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    title character varying(255) NOT NULL,
    descriptrion character varying(4000),
    truth numeric,
    created_by integer,
    argument_pro_id integer,
    reference_id integer
);

CREATE SEQUENCE debates_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;

ALTER SEQUENCE debates_id_seq OWNED BY debates.id;

ALTER TABLE ONLY debates ALTER COLUMN id SET DEFAULT nextval('debates_id_seq'::regclass);

SELECT pg_catalog.setval('debates_id_seq', 1, false);

ALTER TABLE ONLY debates
    ADD CONSTRAINT debates_pkey PRIMARY KEY (id);

ALTER TABLE ONLY debates
    ADD CONSTRAINT debates_user_id_fkey FOREIGN KEY (created_by)
      REFERENCES users (id) MATCH SIMPLE
      ON UPDATE CASCADE ON DELETE SET NULL;

ALTER TABLE ONLY debates
    ADD CONSTRAINT debates_argument_pro_id_fkey FOREIGN KEY (argument_pro_id)
      REFERENCES arguments (id) MATCH SIMPLE
      ON UPDATE CASCADE ON DELETE SET NULL;

ALTER TABLE ONLY debates
    ADD CONSTRAINT debates_argument_con_id_fkey FOREIGN KEY (argument_con_id)
      REFERENCES arguments (id) MATCH SIMPLE
      ON UPDATE CASCADE ON DELETE SET NULL;

ALTER TABLE ONLY debates
    ADD CONSTRAINT debates_references_id_fkey FOREIGN KEY (reference_id)
      REFERENCES 'references' (id) MATCH SIMPLE
      ON UPDATE CASCADE ON DELETE SET NULL;

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back

DROP TABLE debates;
DROP TABLE users;