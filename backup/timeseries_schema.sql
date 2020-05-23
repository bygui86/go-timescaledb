--
-- PostgreSQL database dump
--

-- Dumped from database version 12.3
-- Dumped by pg_dump version 12.3

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- Name: conditions; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.conditions (
    "time" timestamp with time zone NOT NULL,
    location text NOT NULL,
    temperature double precision
);


ALTER TABLE public.conditions OWNER TO postgres;

--
-- Name: conditions_time_idx; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX conditions_time_idx ON public.conditions USING btree ("time" DESC);


--
-- Name: conditions ts_insert_blocker; Type: TRIGGER; Schema: public; Owner: postgres
--



--
-- PostgreSQL database dump complete
--

