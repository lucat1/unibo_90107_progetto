SELECT e.title
FROM events AS e
WHERE ( 
  SELECT count(s.id)
  FROM seats AS s, sectors_events_cost AS sep
  WHERE s.sector = sep.sector AND s.id NOT IN (
    SELECT seat FROM tickets WHERE event = e.id
  ) AND sep.event = e.id
) = 0
