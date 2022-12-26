SELECT SUM(sep.price) - (
  SELECT s.siae_price + s.cachet + SUM(spe.price) + v.price
  FROM shows AS s, events AS e, events_service_providers_serve AS spe, venues AS v
  WHERE e.show = s.id AND spe.event = e.id AND e.venue = v.id AND e.id = ...
) expenses
FROM seats AS s, sectors_events_cost AS sec, tickets AS t
WHERE s.sector = sep.sector AND t.seat = s.id
  AND sep.event = ... AND t.event = ... 
