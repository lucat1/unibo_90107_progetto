SELECT s.row, s.col
FROM seats AS s, sectors_events_cost AS sep
WHERE s.sector = sep.sector AND sep.event = ...