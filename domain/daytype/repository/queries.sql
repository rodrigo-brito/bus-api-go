-- name: by-bus
SELECT DISTINCT d.id, d.name,
  CAST(d.active AS UNSIGNED) as active
FROM `day_type` d
INNER JOIN `schedule` s ON d.id = s.daytype_id
WHERE s.bus_id = ?
ORDER BY d.id;
