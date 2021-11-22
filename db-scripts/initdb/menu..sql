-- Table: public.menu

-- DROP TABLE public.menu;

CREATE TABLE IF NOT EXISTS public.menu
(
    id bigint NOT NULL DEFAULT nextval('menu_id_seq'::regclass),
    name character varying(150) COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    CONSTRAINT menu_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE public.menu
    OWNER to postgres;