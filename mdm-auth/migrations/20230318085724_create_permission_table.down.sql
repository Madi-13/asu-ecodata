BEGIN;

DROP TABLE IF EXISTS mdm_auth.permissions;
DROP INDEX IF EXISTS mdm_auth.permissions_code_role_name_uidx;

COMMIT;