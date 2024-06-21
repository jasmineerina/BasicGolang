-- Write a query to print the ID and StudentGrade for each record in the STUDENT table.
-- Sort the output by student ID, ascending, and use the following format.


-- Student STUDENT.ID has grade: StudentGrade


-- If Score < 20, StudentGrade = F
-- If 20 <= Score < 40, StudentGrade = D
-- If 40 <= Score < 60, StudentGrade = C
-- If 60 <= Score < 80, StudentGrade = B
-- If Score >= 80, StudentGrade = A

SELECT Id, -- pilih id dr tabel students 
CASE -- membuat kondisi
    WHEN Score < 20 THEN 'F' 
    WHEN Score < 40 THEN 'D'
    WHEN Score < 60 THEN 'C'
    WHEN Score < 80 THEN 'B'
    ELSE  'A'
END AS Sudent_Grade -- menutup kondisi dengan menyimpan hasil di student_grade
FROM Students -- menentukan sumber data dr tabel students
ORDER BY id; -- mengurutkan hasil berdasarkan id