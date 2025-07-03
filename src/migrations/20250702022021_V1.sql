-- Drop schema named "public"
-- DROP SCHEMA "public" CASCADE;
-- Rename a column from "id_tag" to "id"
ALTER TABLE "dw"."dim_tag" RENAME COLUMN "id_tag" TO "id";
-- Rename a column from "name_tag" to "name"
ALTER TABLE "dw"."dim_tag" RENAME COLUMN "name_tag" TO "name";
-- Modify "dim_tag" table
ALTER TABLE "dw"."dim_tag" ADD COLUMN "color" character varying(100) NOT NULL;
-- Modify "fato_cards" table
ALTER TABLE "dw"."fato_cards" DROP CONSTRAINT "fato_cards_id_card_fkey", DROP COLUMN "id_time";
-- Rename a column from "id_project" to "id"
ALTER TABLE "dw"."dim_project" RENAME COLUMN "id_project" TO "id";
-- Rename a column from "name_project" to "name"
ALTER TABLE "dw"."dim_project" RENAME COLUMN "name_project" TO "name";
-- Modify "dim_project" table
ALTER TABLE "dw"."dim_project" ADD COLUMN "created_date" timestamp NOT NULL, ADD COLUMN "modified_date" timestamp NOT NULL;
-- Rename a column from "id_status" to "id"
ALTER TABLE "dw"."dim_status" RENAME COLUMN "id_status" TO "id";
-- Rename a column from "name_status" to "name"
ALTER TABLE "dw"."dim_status" RENAME COLUMN "name_status" TO "name";
-- Rename a column from "id_role" to "id"
ALTER TABLE "dw"."dim_role" RENAME COLUMN "id_role" TO "id";
-- Rename a column from "name_role" to "name"
ALTER TABLE "dw"."dim_role" RENAME COLUMN "name_role" TO "name";
-- Rename a column from "id_user" to "id"
ALTER TABLE "dw"."dim_user" RENAME COLUMN "id_user" TO "id";
-- Rename a column from "name_user" to "full_name"
ALTER TABLE "dw"."dim_user" RENAME COLUMN "name_user" TO "full_name";
-- Rename a column from "email" to "color"
ALTER TABLE "dw"."dim_user" RENAME COLUMN "email" TO "color";
-- Modify "dim_user" table
ALTER TABLE "dw"."dim_user" DROP CONSTRAINT "dim_user_id_role_fkey", DROP COLUMN "password", ADD CONSTRAINT "dim_user_id_role_fkey" FOREIGN KEY ("id") REFERENCES "dw"."dim_role" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION;
-- Drop "dim_card" table
DROP TABLE "dw"."dim_card";
-- Drop "dim_time" table
DROP TABLE "dw"."dim_time";
-- Drop "dim_day" table
DROP TABLE "dw"."dim_day";
-- Drop "dim_hour" table
DROP TABLE "dw"."dim_hour";
-- Drop "dim_minute" table
DROP TABLE "dw"."dim_minute";
-- Drop "dim_month" table
DROP TABLE "dw"."dim_month";
-- Drop "dim_year" table
DROP TABLE "dw"."dim_year";
