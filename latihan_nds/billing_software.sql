-- You are working on a billing application and need to get a list of customers who have a positive balance.
-- The result should have the following columns: iban | amount.
-- The result should be sorted in descending order by amount.

-- Note:
-- • Only customers whose balance is positive should be included in the report.
-- • For MS SQL Server users - in order to match the required format, 
--   you can use FORMAT(amount, 'N') in your query.

SELECT FROM iban, FORMAT(amount, 'N') AS amount --- amount diformat sebagai angka
FROM customers --- kolom iban dan amount dr tabel customers
WHERE amount > 0 --- filter data amount yang berisi angka positif
ORDER BY amount DESC; --- mengurutkan hasil berdasarkan amount (dr besar ke kecil)