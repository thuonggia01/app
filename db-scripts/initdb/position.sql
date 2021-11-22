-- Table: public.position

-- DROP TABLE public."position";

CREATE TABLE IF NOT EXISTS public."position"
(
    id bigint NOT NULL DEFAULT nextval('position_id_seq'::regclass),
    name character varying COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    CONSTRAINT position_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE public."position"
    OWNER to postgres;