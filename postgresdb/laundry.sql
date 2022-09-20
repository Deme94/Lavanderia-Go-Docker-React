--
-- PostgreSQL database dump
--

-- Dumped from database version 14.3 (Ubuntu 14.3-1.pgdg20.04+1)
-- Dumped by pg_dump version 14.3 (Ubuntu 14.3-1.pgdg20.04+1)

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
-- Name: users; Type: TABLE; Schema: public; Owner: deme
--

CREATE TABLE public.users (
    id integer NOT NULL,
    email character varying(255) NOT NULL,
    password character varying(80) NOT NULL,
    created_on timestamp without time zone NOT NULL,
    last_login timestamp without time zone
);


ALTER TABLE public.users OWNER TO deme;

--
-- Name: users_user_id_seq; Type: SEQUENCE; Schema: public; Owner: deme
--

CREATE SEQUENCE public.users_user_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.users_user_id_seq OWNER TO deme;

--
-- Name: users_user_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: deme
--

ALTER SEQUENCE public.users_user_id_seq OWNED BY public.users.id;


--
-- Name: users id; Type: DEFAULT; Schema: public; Owner: deme
--

ALTER TABLE ONLY public.users ALTER COLUMN id SET DEFAULT nextval('public.users_user_id_seq'::regclass);


--
-- Data for Name: users; Type: TABLE DATA; Schema: public; Owner: deme
--

COPY public.users (id, email, password, created_on, last_login) FROM stdin;
4	pruebaTest@gmail.com	$2a$12$AJXCyIXF7vxiCYagP3pHIOj46Xqs0KfSkCZrFMgB/ejhsLM2P6CXS	2022-05-02 15:32:00.454697	\N
5	pruebatest2@gmail.com	$2a$12$V6hIe9ciFS1H25bJRmaiE.B4Uk5JWYwKrugsVQsbLdCzX9kAZE7AC	2022-05-02 15:40:26.28089	\N
6	pruebastest2@gmail.com	$2a$12$xH0mZPkyiJ6ZB1/AioOW8urf3w4Z9iY8eQAfim9RwvdHpyZESDhnW	2022-05-02 15:42:09.362553	\N
7	pruebatest3@gmail.com	$2a$12$n5giFl0dMSh5GrY.bK.WHu2El1l4I.tpw/kRYrIr.h5GwlOXOiodS	2022-05-02 15:46:03.565257	\N
8	pruebatest4@gmail.com	$2a$12$cQFODI/twAQGbthh2IJOXuLZ.0X04XhhStcETkH2FVuSxevwjLL3.	2022-05-02 15:47:12.708445	\N
9	pruebatest5@gmail.com	$2a$12$06OpBu8hvefvJdYKidM7IeBDtIDMr1VCh9R2vyfDlT6TuN5mp0T3u	2022-05-02 15:54:44.79466	\N
10	pruebatest6@gmail.com	$2a$12$3c9C6NuLE0S6KwsAYhZrmOGQSPbPrHE3g9EGqCBmLAqyONdD1P.ki	2022-05-02 15:58:24.017494	\N
11	pruebatest8@gmail.com	$2a$12$Bear9M6ZBFoYg6/0RE6MQu50HtAUg8zYxg7Lg.yoFvG0tcbLnl3TW	2022-05-02 16:01:40.110668	\N
12	pruebatest9@gmail.com	$2a$12$7OSqIYIH7R8GF8GiPTr3euf6Z9a/HSmrTsT4/DE64MX1AMKEV5Uwu	2022-05-02 16:01:51.742896	\N
13	pruebatest10@gmail.com	$2a$12$mD52.LiNPvu12EP0wcRq4ugQ39ZmIOXAgwt..mopWEb.Lj6fUw1zi	2022-05-02 16:10:43.239803	\N
2	deme1994@gmail.com	$2a$12$zD58C1N70giiPqX.Qa0QmuaY/2nXBPgcqy4dLP5gdh53zx.q4VlsW	2022-05-01 20:29:12.772275	2022-05-02 19:48:12.133143
\.


--
-- Name: users_user_id_seq; Type: SEQUENCE SET; Schema: public; Owner: deme
--

SELECT pg_catalog.setval('public.users_user_id_seq', 13, true);


--
-- Name: users users_email_key; Type: CONSTRAINT; Schema: public; Owner: deme
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_email_key UNIQUE (email);


--
-- Name: users users_pkey; Type: CONSTRAINT; Schema: public; Owner: deme
--

ALTER TABLE ONLY public.users
    ADD CONSTRAINT users_pkey PRIMARY KEY (id);


--
-- PostgreSQL database dump complete
--

