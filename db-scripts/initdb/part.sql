-- Table: public.part

-- DROP TABLE public.part;

CREATE TABLE IF NOT EXISTS public.part
(
    id bigint NOT NULL DEFAULT nextval('part_id_seq'::regclass),
    name character varying(150) COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    CONSTRAINT part_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE public.part
    OWNER to postgres;