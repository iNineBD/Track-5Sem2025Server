-- Modify "dim_project" table
ALTER TABLE "dw"."dim_project" DROP COLUMN "id_platform";
-- Drop "dim_platform" table
DROP TABLE "dw"."dim_platform";
