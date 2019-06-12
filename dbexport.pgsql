--
-- PostgreSQL database dump
--

-- Dumped from database version 11.3
-- Dumped by pg_dump version 11.3

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

SET default_with_oids = false;

--
-- Name: incidents; Type: TABLE; Schema: public; Owner: jainam
--

CREATE TABLE public.incidents (
    id bigint NOT NULL,
    alertname character varying(255),
    starts_at character varying(255),
    address character varying(255),
    status character varying(255)
);


ALTER TABLE public.incidents OWNER TO jainam;

--
-- Name: incidents_id_seq; Type: SEQUENCE; Schema: public; Owner: jainam
--

ALTER TABLE public.incidents ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.incidents_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: mapping; Type: TABLE; Schema: public; Owner: jainam
--

CREATE TABLE public.mapping (
    alert_type character varying,
    script character varying
);


ALTER TABLE public.mapping OWNER TO jainam;

--
-- Name: operations; Type: TABLE; Schema: public; Owner: jainam
--

CREATE TABLE public.operations (
    id bigint NOT NULL,
    surgeon_id character varying(255),
    script character varying(255),
    status character varying(255),
    alert_id character varying(255)
);


ALTER TABLE public.operations OWNER TO jainam;

--
-- Name: operations_id_seq; Type: SEQUENCE; Schema: public; Owner: jainam
--

ALTER TABLE public.operations ALTER COLUMN id ADD GENERATED ALWAYS AS IDENTITY (
    SEQUENCE NAME public.operations_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1
);


--
-- Name: schema_migrations; Type: TABLE; Schema: public; Owner: jainam
--

CREATE TABLE public.schema_migrations (
    version bigint NOT NULL,
    dirty boolean NOT NULL
);


ALTER TABLE public.schema_migrations OWNER TO jainam;

--
-- Name: incidents id; Type: CONSTRAINT; Schema: public; Owner: jainam
--

ALTER TABLE ONLY public.incidents
    ADD CONSTRAINT id PRIMARY KEY (id);


--
-- Name: operations operations_pkey; Type: CONSTRAINT; Schema: public; Owner: jainam
--

ALTER TABLE ONLY public.operations
    ADD CONSTRAINT operations_pkey PRIMARY KEY (id);


--
-- Name: schema_migrations schema_migrations_pkey; Type: CONSTRAINT; Schema: public; Owner: jainam
--

ALTER TABLE ONLY public.schema_migrations
    ADD CONSTRAINT schema_migrations_pkey PRIMARY KEY (version);


--
-- PostgreSQL database dump complete
--

