-- Create "bfi_reports" table
CREATE TABLE "bfi_reports" ("oid" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "conscientiousness" jsonb NOT NULL, "extraversion" jsonb NOT NULL, "agreeableness" jsonb NOT NULL, "neuroticism" jsonb NOT NULL, "openness" jsonb NOT NULL, PRIMARY KEY ("oid"));
-- Create "users" table
CREATE TABLE "users" ("oid" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "github_username" character varying NOT NULL, "univeresity_id" character varying NULL, "role" character varying NOT NULL DEFAULT 'student', "bfi_report_student" uuid NULL, PRIMARY KEY ("oid"), CONSTRAINT "users_bfi_reports_student" FOREIGN KEY ("bfi_report_student") REFERENCES "bfi_reports" ("oid") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "users_bfi_report_student_key" to table: "users"
CREATE UNIQUE INDEX "users_bfi_report_student_key" ON "users" ("bfi_report_student");
-- Create index "users_github_username_key" to table: "users"
CREATE UNIQUE INDEX "users_github_username_key" ON "users" ("github_username");
-- Create index "users_univeresity_id_key" to table: "users"
CREATE UNIQUE INDEX "users_univeresity_id_key" ON "users" ("univeresity_id");
-- Create "courses" table
CREATE TABLE "courses" ("oid" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "name" character varying NOT NULL, "code" character varying NOT NULL, "colour" character varying NOT NULL DEFAULT '#000000', "user_created_courses" uuid NULL, PRIMARY KEY ("oid"), CONSTRAINT "courses_users_created_courses" FOREIGN KEY ("user_created_courses") REFERENCES "users" ("oid") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create index "courses_code_key" to table: "courses"
CREATE UNIQUE INDEX "courses_code_key" ON "courses" ("code");
-- Create "groups" table
CREATE TABLE "groups" ("oid" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "course_groups" uuid NOT NULL, "group_created_by" uuid NOT NULL, PRIMARY KEY ("oid"), CONSTRAINT "groups_courses_groups" FOREIGN KEY ("course_groups") REFERENCES "courses" ("oid") ON UPDATE NO ACTION ON DELETE NO ACTION, CONSTRAINT "groups_users_created_by" FOREIGN KEY ("group_created_by") REFERENCES "users" ("oid") ON UPDATE NO ACTION ON DELETE NO ACTION);
-- Create "group_students" table
CREATE TABLE "group_students" ("group_id" uuid NOT NULL, "user_id" uuid NOT NULL, PRIMARY KEY ("group_id", "user_id"), CONSTRAINT "group_students_group_id" FOREIGN KEY ("group_id") REFERENCES "groups" ("oid") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "group_students_user_id" FOREIGN KEY ("user_id") REFERENCES "users" ("oid") ON UPDATE NO ACTION ON DELETE CASCADE);
-- Create "bfi_questions" table
CREATE TABLE "bfi_questions" ("oid" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "question_slo" character varying NOT NULL, "question_eng" character varying NOT NULL, "facet" character varying NOT NULL, "dimension" character varying NOT NULL, "influence" character varying NOT NULL, "alpha" double precision NOT NULL, PRIMARY KEY ("oid"));
-- Create "bfi_answers" table
CREATE TABLE "bfi_answers" ("oid" uuid NOT NULL, "created_at" timestamptz NOT NULL, "updated_at" timestamptz NOT NULL, "value" bigint NULL, "bfi_question_bfi_answers" uuid NULL, "user_bfi_answers" uuid NULL, PRIMARY KEY ("oid"), CONSTRAINT "bfi_answers_bfi_questions_bfi_answers" FOREIGN KEY ("bfi_question_bfi_answers") REFERENCES "bfi_questions" ("oid") ON UPDATE NO ACTION ON DELETE SET NULL, CONSTRAINT "bfi_answers_users_bfi_answers" FOREIGN KEY ("user_bfi_answers") REFERENCES "users" ("oid") ON UPDATE NO ACTION ON DELETE SET NULL);
-- Create "course_students" table
CREATE TABLE "course_students" ("course_id" uuid NOT NULL, "user_id" uuid NOT NULL, PRIMARY KEY ("course_id", "user_id"), CONSTRAINT "course_students_course_id" FOREIGN KEY ("course_id") REFERENCES "courses" ("oid") ON UPDATE NO ACTION ON DELETE CASCADE, CONSTRAINT "course_students_user_id" FOREIGN KEY ("user_id") REFERENCES "users" ("oid") ON UPDATE NO ACTION ON DELETE CASCADE);
