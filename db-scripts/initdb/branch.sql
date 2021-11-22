CREATE TABLE IF NOT EXISTS public.branch
(
    id bigint NOT NULL DEFAULT nextval('branch_id_seq'::regclass),
    name character varying(150) COLLATE pg_catalog."default" NOT NULL,
    created_at timestamp without time zone,
    updated_at timestamp without time zone,
    deleted_at timestamp without time zone,
    CONSTRAINT branch_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE public.branch
    OWNER to postgres;