SELECT SUM(sep.price)
FROM seats AS s, sectors_events_cost AS sec, tickets AS t
WHERE s.sector = sep.sector AND t.seat = s.id
  AND sep.event = ... AND t.event = ... 

