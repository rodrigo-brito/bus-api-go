-- name: by-bus
SELECT id, origin, destiny, observation, time
FROM schedule
WHERE bus_id = ?;

-- name: by-bus-daytype
SELECT id, origin, destiny, observation, time
FROM schedule
WHERE bus_id = ? and daytype_id = ?;