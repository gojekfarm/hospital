CREATE TABLE operations
(
    id bigint NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 9223372036854775807 CACHE 1 ),
    application_id character varying(255) COLLATE pg_catalog."default",
    script character varying(255) COLLATE pg_catalog."default",
    status character varying(255) COLLATE pg_catalog."default",
    alert_id character varying(255) COLLATE pg_catalog."default",
    CONSTRAINT operations_pkey PRIMARY KEY (id)
);
