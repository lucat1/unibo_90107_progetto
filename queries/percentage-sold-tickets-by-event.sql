SELECT (
  SELECT COUNT(*)
  FROM tickets
  WHERE event = ...
) * 100.0 / SUM(s.capacity)
FROM sectors_events_cost AS sec, sectors AS s
WHERE s.id = sec.sector AND sec.event = ... 
