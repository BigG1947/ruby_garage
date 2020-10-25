-- 1. get all statuses, not repeating, alphabetically ordered
SELECT DISTINCT status FROM tasks ORDER BY status;
-- 2. get the count of all tasks in each project, order by tasks count descending
SELECT COUNT(id) AS task_count FROM tasks GROUP BY project_id ORDER BY task_count DESC;
-- 3. get the count of all tasks in each project, order by projects names
SELECT COUNT(t.id) AS task_count, p.name FROM tasks t INNER JOIN projects p ON p.id = t.project_id GROUP BY project_id ORDER BY p.name;
-- 4 get the tasks for all projects having the name beginning with "N" letter
SELECT id, name, status, project_id FROM tasks WHERE name LIKE 'n%';
-- 5 get the list of all projects containing the 'a' letter in the middle of the name, and show the tasks count near each project. Mention that there can exist projects without tasks and tasks with project_id = NULL
SELECT p.name, COUNT(t.id) FROM projects p LEFT JOIN tasks t ON p.id = t.project_id GROUP BY p.id
-- 6 get the list of tasks with duplicate names. Order alphabetically
-- variant #1
SELECT t.* FROM tasks t WHERE t.name IN (SELECT name FROM tasks GROUP BY name HAVING COUNT(id) > 1) ORDER BY t.name;
-- variant #2
SELECT name, COUNT (name) FROM tasks GROUP BY name HAVING COUNT(name) > 1 ORDER BY name;
-- 7 get list of tasks having several exact matches of both name and status, from the project 'Garage'. Order by matches count
-- variant #1
SELECT t.*, p.name FROM tasks t INNER JOIN projects p on t.project_id = p.id WHERE p.name = 'Garage' AND (t.name, t.status) IN (SELECT name, status FROM tasks GROUP BY name, status HAVING COUNT(id) > 1)
-- variant #2
SELECT t.name, t.status, p.name, COUNT(t.id) AS count_task FROM tasks t INNER JOIN projects p ON t.project_id = p.id WHERE p.name = 'Garage' AND (t.name, t.status) IN (SELECT name, status FROM tasks GROUP BY name, status HAVING COUNT(id) > 1) GROUP BY t.name, t.status ORDER BY count_task, t.name, t.status
-- 8 get the list of project names having more than 10 tasks in status 'completed'. Order by project_id
SELECT p.name, COUNT(t.id) AS count_task FROM projects p INNER JOIN tasks t on p.id = t.project_id WHERE t.status = 'completed' GROUP BY p.name HAVING count_task > 10 ORDER BY p.id