-- Create "users" table
CREATE TABLE "users" ("id" uuid NOT NULL, "full_name" character varying NOT NULL, "bio" character varying NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NULL, PRIMARY KEY ("id"));
-- Create "accounts" table
CREATE TABLE "accounts" ("id" uuid NOT NULL, "email" character varying NOT NULL, "password" character varying NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "deleted_at" timestamptz NULL, "user_account" uuid NULL, PRIMARY KEY ("id"), CONSTRAINT "accounts_users_account" FOREIGN KEY ("user_account") REFERENCES "users" ("id") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "accounts_email_key" to table: "accounts"
CREATE UNIQUE INDEX "accounts_email_key" ON "accounts" ("email");
-- Create index "accounts_user_account_key" to table: "accounts"
CREATE UNIQUE INDEX "accounts_user_account_key" ON "accounts" ("user_account");
