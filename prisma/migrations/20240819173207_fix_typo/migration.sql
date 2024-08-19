/*
  Warnings:

  - You are about to drop the column `t0_account` on the `transaction` table. All the data in the column will be lost.
  - Added the required column `to_account` to the `transaction` table without a default value. This is not possible if the table is not empty.

*/
-- AlterTable
ALTER TABLE "transaction" DROP COLUMN "t0_account",
ADD COLUMN     "to_account" VARCHAR(100) NOT NULL;
