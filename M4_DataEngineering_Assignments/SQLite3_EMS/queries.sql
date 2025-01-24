--Q1. List the names of employees hired after January 1, 2021.
SELECT name 
FROM Employees 
WHERE hire_date > '2021-01-01';

--Q2. Calculate the average salary of employees in each department.
SELECT d.name AS department_name, 
       AVG(e.salary) AS average_salary
FROM Employees e
JOIN Departments d ON e.department_id = d.department_id
GROUP BY d.department_id;

--Q3. Find the department name where the total salary is the highest.
SELECT d.name AS department_name
FROM Employees e
JOIN Departments d ON e.department_id = d.department_id
GROUP BY d.department_id
ORDER BY SUM(e.salary) DESC
LIMIT 1;

--Q4. List all departments that currently have no employees assigned.
SELECT name 
FROM Departments d
WHERE NOT EXISTS (
    SELECT 1 
    FROM Employees e 
    WHERE e.department_id = d.department_id
);

--Q5. Fetch all employee details along with their department names.
SELECT e.*, d.name AS department_name
FROM Employees e
JOIN Departments d ON e.department_id = d.department_id;
