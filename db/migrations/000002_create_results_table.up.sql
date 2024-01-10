-- Initial public schema relates to Library 0.x

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SET client_min_messages = warning;
SET row_security = off;

CREATE EXTENSION IF NOT EXISTS plpgsql WITH SCHEMA pg_catalog;
CREATE EXTENSION IF NOT EXISTS "uuid-ossp" WITH SCHEMA pg_catalog;

SET search_path = public, pg_catalog;
SET default_tablespace = '';

-- results
CREATE TABLE results (
    id uuid NOT NULL DEFAULT uuid_generate_v1mc(),
    runner_id uuid NOT NULL,
    race_result interval NOT NULL,
    location text NOT NULL,
    position integer,
    year integer NOT NULL,
    CONSTRAINT results_pk PRIMARY KEY (id),
    CONSTRAINT fk_results_runner_id FOREIGN KEY (runner_id)
        REFERENCES runners (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
);