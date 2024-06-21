-- As part of a smart home application, create a query that, based on data from meter readings, 
-- calculates the total electricity consumption, the name of the most expensive tariff consumed, 
-- and the total cost for invoicing.

-- The result should have the following columns: 
-- username | email | highest tariff | consumption | total_cost.

-- • username - account username
-- • email - account email address
-- • highest tariff - name of the most expensive tariff that is consumed
-- • consumption - total consumption amount
-- • total_cost - total cost of all consumed electricity, rounded to two decimal places
-- The result should be sorted in ascending order by username.

-- Note:
-- • The total cost of all electricity consumed is the sum of the products of all meter readings and their respective consumed tariffs

SELECT
    u.username, -- ambil dr tabel user
    u.email,
    t.name AS highest_tariff, -- ambil nama tarif tertinggi yg dikenakan pengguna
    SUM(mr.reading) AS consumption, -- menjumlahkan semua meter reading tiap pengguna
    ROUND(SUM(mr.reading * tr.price), 2) AS total_cost -- hitung total biaya pengunaan
FROM
    users u
JOIN
    meter_readings mr ON u.user_id = mr.user_id -- gabung tabel
JOIN
    tariffs tr ON mr.tariff_id = tr.tariff_id
JOIN
    (
        SELECT
            user_id,
            MAX(price) as max_price -- untuk cari harga tarif max tiap pengguna dr tabel tariffs
        FROM
            tariffs
        GROUP BY
            user_id
    )max_tariff ON tr.user_id = max_tariff.user_id AND tr.price = max_tariff.max_price
    GROUP BY
        u.username, u.email, t.name -- mengelompokkan hasil berdasarkan kolom
    ORDER BY
        u.username ASC; -- mengurutkan hasil berdasarkan usn dr kecil ke besar