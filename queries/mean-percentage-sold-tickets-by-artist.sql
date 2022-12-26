SELECT AVG((
  SELECT COUNT(*)
  FROM tickets
  WHERE event = e.id
) * 100.0 / SUM(s.capacity))
FROM sectors_events_cost AS sec, 
   sectors AS s,
   show    AS sh, 
   events  AS e
WHERE s.id = sec.sector AND 
  sec.event = e.id      AND 
  sh.artist = ...       AND 
  sh.id = e.show
