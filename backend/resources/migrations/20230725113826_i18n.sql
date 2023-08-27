-- Modify "bfi_questions" table
ALTER TABLE "bfi_questions" DROP COLUMN "question_slo", DROP COLUMN "question_eng", ADD COLUMN "questiono" character varying NOT NULL;
