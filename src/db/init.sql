CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

DROP TABLE IF EXISTS users CASCADE;

CREATE TABLE public.users
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    firstname character varying(255) COLLATE pg_catalog."default" NOT NULL,
    lastname character varying(255) COLLATE pg_catalog."default" NOT NULL,
    email character varying(255) COLLATE pg_catalog."default" NOT NULL,
    createdon character varying(255) COLLATE pg_catalog."default" NULL,
    updatedon character varying(255) COLLATE pg_catalog."default" NULL,
    deletedon character varying(255) COLLATE pg_catalog."default" NULL,
    CONSTRAINT users_pkey PRIMARY KEY (id)
);

ALTER TABLE public.users OWNER to postgres;

DROP TABLE IF EXISTS datauploads;

CREATE TABLE public.datauploads
(
    id uuid NOT NULL DEFAULT uuid_generate_v4(),
    userid uuid REFERENCES users (id),
    createdon timestamp NULL,
    updatedon timestamp NULL,
    deletedon timestamp NULL,
    updatedby uuid REFERENCES users (id) NULL,
    CONSTRAINT datauploads_pkey PRIMARY KEY (id)
);

ALTER TABLE public.datauploads OWNER to postgres;

INSERT INTO users (firstname,lastname, email) VALUES ('George','Washington', 'gw@whitehouse.com');
INSERT INTO users (firstname,lastname, email) VALUES ('Penelope','Penultimate', 'pp@there.com');
INSERT INTO users (firstname,lastname, email) VALUES ('Magnus','Carlson', 'mc@chess.com');
