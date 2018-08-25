-- +goose Up
CREATE TABLE "users" (
  "id" bigserial primary key,
  "email" character varying(255) NOT NULL,
  "name" character varying(255) NOT NULL,
  "image_url" text NOT NULL,
  "role" integer DEFAULT 0 NOT NULL,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL
);

CREATE UNIQUE INDEX  "index_users_on_email" ON "users" USING btree ("email");

CREATE TABLE "notifications" (
  "id" bigserial primary key,
  "name" character varying(255) NOT NULL,
  "service" integer NOT NULL,
  "webhook_url" text,
  "channel" text,
  "user_name" text,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL
);

CREATE TABLE "workspaces" (
  "id" bigserial primary key,
  "name" character varying(255) NOT NULL,
  "notification_type" integer NOT NULL DEFAULT 0,
  "notification_id" bigint REFERENCES notifications(id) ON DELETE SET NULL,
  "description" text,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL
);

INSERT INTO workspaces (name, created_at, updated_at) VALUES ('default workspace', now(), now());

CREATE TABLE "posts" (
  "id" bigserial primary key,
  "workspace_id" bigint NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
  "title" character varying(255) NOT NULL,
  "body" text NOT NULL,
  "creator_user_id" bigint REFERENCES users(id) ON DELETE SET NULL,
  "editor_user_id" bigint REFERENCES users(id) ON DELETE SET NULL,
  "edit_status" integer NOT NULL,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL
);

CREATE INDEX  "index_posts_on_created_at" ON "posts" USING btree ("workspace_id", "created_at");
CREATE INDEX  "index_posts_on_updated_at" ON "posts" USING btree ("workspace_id", "updated_at");

CREATE TABLE "post_histories" (
  "id" bigserial primary key,
  "version" bigint NOT NULL,
  "post_id" bigint NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
  "title" character varying(255) NOT NULL,
  "body" text NOT NULL,
  "editor_user_id" bigint REFERENCES users(id) ON DELETE SET NULL,
  "edit_status" integer NOT NULL,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL
);

CREATE UNIQUE INDEX  "index_post_histories_on_post_id_version" ON "post_histories" USING btree ("post_id", "version");

CREATE TABLE "user_tokens" (
  "id" bigserial primary key,
  "user_id" bigint NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  "provider" integer NOT NULL,
  "token_json" text,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL
);

CREATE UNIQUE INDEX  "index_user_tokens_on_user_id_provider" ON "user_tokens" USING btree ("user_id", "provider");

CREATE TABLE "user_signin_histories" (
  "id" bigserial primary key,
  "user_id" bigint NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  "provider" integer NOT NULL,
  "ipv4_address" character varying(32),
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL
);

CREATE TABLE "groups" (
  "id" bigserial primary key,
  "name" character varying(255) NOT NULL,
  "notification_id" bigint REFERENCES notifications(id) ON DELETE SET NULL,
  "description" text,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL
);

INSERT INTO groups (name, created_at, updated_at) VALUES ('default group', now(), now());

CREATE TABLE "group_users" (
  "id" bigserial primary key,
  "user_id" bigint NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  "group_id" bigint NOT NULL REFERENCES groups(id) ON DELETE CASCADE,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL
);

CREATE UNIQUE INDEX  "index_group_users_on_user_id_group_id" ON "group_users" USING btree ("user_id", "group_id");

CREATE TABLE "roles" (
  "id" bigserial primary key,
  "name" character varying(255) NOT NULL,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL
);

INSERT INTO roles (name, created_at, updated_at) VALUES ('admin', now(), now());

CREATE TABLE "role_permissions" (
  "id" bigserial primary key,
  "role_id" bigint NOT NULL REFERENCES roles(id) ON DELETE CASCADE,
  "permission" integer NOT NULL,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL
);

CREATE UNIQUE INDEX  "index_role_permissions_on_role_id_permission" ON "role_permissions" USING btree ("role_id", "permission");

CREATE TABLE "post_comments" (
  "id" bigserial primary key,
  "post_id" bigint NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
  "comment" text NOT NULL,
  "user_id" bigint REFERENCES users(id) ON DELETE SET NULL,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL
);

CREATE INDEX  "index_post_comments_on_post_id_created_at" ON "post_comments" USING btree ("post_id", "created_at");

CREATE TABLE "tags" (
  "id" bigserial primary key,
  "name" character varying(255) NOT NULL,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL
);

CREATE TABLE "posts_tags" (
  "id" bigserial primary key,
  "post_id" bigint NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
  "tag_id" bigint NOT NULL REFERENCES tags(id) ON DELETE CASCADE,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL
);

CREATE UNIQUE INDEX  "index_posts_tags_on_post_id_tag_id" ON "posts_tags" USING btree ("post_id", "tag_id");

CREATE TABLE "post_templates" (
  "id" bigserial primary key,
  "workspace_id" bigint NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
  "post_id" bigint NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL
);

CREATE UNIQUE INDEX  "index_post_templates_on_post_id" ON "post_templates" USING btree ("post_id");

CREATE TABLE "post_likes" (
  "id" bigserial primary key,
  "post_id" bigint NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
  "user_id" bigint NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL
);

CREATE UNIQUE INDEX  "index_post_likes_on_post_id_user_id" ON "post_likes" USING btree ("post_id", "user_id");

CREATE TABLE "post_reads" (
  "id" bigserial primary key,
  "post_id" bigint NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
  "user_id" bigint NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL
);

CREATE UNIQUE INDEX  "index_post_reads_on_post_id_user_id" ON "post_reads" USING btree ("post_id", "user_id");

CREATE TABLE "menu_items" (
  "id" bigserial primary key,
  "workspace_id" bigint NOT NULL REFERENCES workspaces(id) ON DELETE CASCADE,
  "post_id" bigint NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
  "parent_menu_id" bigint REFERENCES menu_items(id) ON DELETE CASCADE,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL
);

CREATE UNIQUE INDEX  "index_menu_items_on_post_id" ON "menu_items" USING btree ("post_id");

CREATE TABLE "post_upload_files" (
  "id" bigserial primary key,
  "post_id" bigint NOT NULL REFERENCES posts(id) ON DELETE CASCADE,
  "file_name" text,
  "file_url" text,
  "user_id" bigint NOT NULL REFERENCES users(id) ON DELETE CASCADE,
  "created_at" timestamp(6) NOT NULL,
  "updated_at" timestamp(6) NOT NULL
);

-- +goose Down
DROP TABLE "post_upload_files";
DROP TABLE "user_tokens";
DROP TABLE "group_users";
DROP TABLE "role_permissions";
DROP TABLE "posts_tags";
DROP TABLE "post_comments";
DROP TABLE "tags";
DROP TABLE "post_templates";
DROP TABLE "post_likes";
DROP TABLE "post_reads";
DROP TABLE "menu_items";
DROP TABLE "groups";
DROP TABLE "roles";
DROP TABLE "post_histories";
DROP TABLE "posts";
DROP TABLE "workspaces";
DROP TABLE "notifications";
DROP TABLE "user_signin_histories";
DROP TABLE "users";
