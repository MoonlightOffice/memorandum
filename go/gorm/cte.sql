CREATE TABLE p (
    id INT PRIMARY KEY,
    name TEXT NOT NULL
);

INSERT INTO p VALUES (1, 'p1');
INSERT INTO p VALUES (2, 'p2');
INSERT INTO p VALUES (3, 'p2');

WITH cte AS (
    SELECT 
        id, 
        name,
        ROW_NUMBER() OVER (PARTITION BY name ORDER BY id) AS rn
    FROM p
)
DELETE FROM p
WHERE id IN (
    SELECT id
    FROM cte
    WHERE rn > 1
);

SELECT 
    id, 
    name,
    ROW_NUMBER() OVER (PARTITION BY name, age) AS rn
FROM p LIMIT 100;

--

WITH RECURSIVE ordered_queue AS (
    -- Anchor: Start from the head node (e.g. known ID)
    SELECT
        q.id,
        q.message,
        q.next_id,
        q.previous_id,
        1 AS depth
    FROM queue q
    WHERE q.previous_id IS NULL  -- Assuming this is the head of the list

    UNION ALL

    -- Recursive part: join on next_id to follow the linked list
    SELECT
        q.id,
        q.message,
        q.next_id,
        q.previous_id,
        oq.depth + 1
    FROM queue q
    JOIN ordered_queue oq ON q.id = oq.next_id
)
SELECT * FROM ordered_queue
ORDER BY depth
LIMIT 100;
