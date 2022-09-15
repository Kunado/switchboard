CREATE TABLE "public"."cname_records" (
    "id" serial NOT NULL,
    "host" character varying(255) NOT NULL,
    "value" character varying(255) NOT NULL,
    "profile_id" integer NOT NULL,
    PRIMARY KEY ("id")
);

ALTER TABLE ONLY "public"."cname_records" ADD CONSTRAINT "cname_records_profile_id_fkey" FOREIGN KEY ("profile_id") REFERENCES "public"."profiles" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;

CREATE TABLE "public"."profiles" (
    "id" serial NOT NULL,
    "name" character varying(255) NOT NULL,
    "enabled" boolean NOT NULL DEFAULT false,
    PRIMARY KEY ("id")
);

ALTER TABLE public.profiles ADD CONSTRAINT uk UNIQUE (name);
