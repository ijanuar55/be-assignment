/*
  Warnings:

  - A unique constraint covering the columns `[account_number]` on the table `account` will be added. If there are existing duplicate values, this will fail.

*/
-- CreateTable
CREATE TABLE "transaction" (
    "id" VARCHAR(100) NOT NULL DEFAULT gen_random_uuid(),
    "from_account" VARCHAR(100) NOT NULL,
    "t0_account" VARCHAR(100) NOT NULL,
    "created_at" TIMESTAMP(3) DEFAULT CURRENT_TIMESTAMP,
    "amount" DOUBLE PRECISION NOT NULL,

    CONSTRAINT "transaction_pkey" PRIMARY KEY ("id")
);

-- CreateIndex
CREATE UNIQUE INDEX "transaction_id_key" ON "transaction"("id");

-- CreateIndex
CREATE UNIQUE INDEX "account_account_number_key" ON "account"("account_number");
