--
-- PostgreSQL database dump
--

-- Dumped from database version 14.5 (Debian 14.5-1.pgdg110+1)
-- Dumped by pg_dump version 14.5 (Homebrew)

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
-- Name: cname_records; Type: TABLE; Schema: public; Owner: testuser
--

CREATE TABLE public.cname_records (
    id integer NOT NULL,
    host character varying(255) NOT NULL,
    value character varying(255) NOT NULL,
    profile_id integer NOT NULL
);


ALTER TABLE public.cname_records OWNER TO testuser;

--
-- Name: cname_records_id_seq; Type: SEQUENCE; Schema: public; Owner: testuser
--

CREATE SEQUENCE public.cname_records_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.cname_records_id_seq OWNER TO testuser;

--
-- Name: cname_records_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: testuser
--

ALTER SEQUENCE public.cname_records_id_seq OWNED BY public.cname_records.id;


--
-- Name: profiles; Type: TABLE; Schema: public; Owner: testuser
--

CREATE TABLE public.profiles (
    id integer NOT NULL,
    name character varying(255) NOT NULL,
    enabled boolean DEFAULT false NOT NULL
);


ALTER TABLE public.profiles OWNER TO testuser;

--
-- Name: profiles_id_seq; Type: SEQUENCE; Schema: public; Owner: testuser
--

CREATE SEQUENCE public.profiles_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.profiles_id_seq OWNER TO testuser;

--
-- Name: profiles_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: testuser
--

ALTER SEQUENCE public.profiles_id_seq OWNED BY public.profiles.id;


--
-- Name: cname_records id; Type: DEFAULT; Schema: public; Owner: testuser
--

ALTER TABLE ONLY public.cname_records ALTER COLUMN id SET DEFAULT nextval('public.cname_records_id_seq'::regclass);


--
-- Name: profiles id; Type: DEFAULT; Schema: public; Owner: testuser
--

ALTER TABLE ONLY public.profiles ALTER COLUMN id SET DEFAULT nextval('public.profiles_id_seq'::regclass);


--
-- Data for Name: cname_records; Type: TABLE DATA; Schema: public; Owner: testuser
--

COPY public.cname_records (id, host, value, profile_id) FROM stdin;
1	localhost	localdns.service	1
\.


--
-- Data for Name: profiles; Type: TABLE DATA; Schema: public; Owner: testuser
--

COPY public.profiles (id, name, enabled) FROM stdin;
1	defualt	t
\.


--
-- Name: cname_records_id_seq; Type: SEQUENCE SET; Schema: public; Owner: testuser
--

SELECT pg_catalog.setval('public.cname_records_id_seq', 1, true);


--
-- Name: profiles_id_seq; Type: SEQUENCE SET; Schema: public; Owner: testuser
--

SELECT pg_catalog.setval('public.profiles_id_seq', 1, true);


--
-- Name: cname_records cname_records_pkey; Type: CONSTRAINT; Schema: public; Owner: testuser
--

ALTER TABLE ONLY public.cname_records
    ADD CONSTRAINT cname_records_pkey PRIMARY KEY (id);


--
-- Name: profiles profiles_pkey; Type: CONSTRAINT; Schema: public; Owner: testuser
--

ALTER TABLE ONLY public.profiles
    ADD CONSTRAINT profiles_pkey PRIMARY KEY (id);


--
-- Name: cname_records cname_records_profile_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: testuser
--

ALTER TABLE ONLY public.cname_records
    ADD CONSTRAINT cname_records_profile_id_fkey FOREIGN KEY (profile_id) REFERENCES public.profiles(id);


--
-- PostgreSQL database dump complete
--

