SELECT t_sh.title, (
  SELECT SUM (
    SELECT SUM(sep.price) - (
      SELECT s.siae_price + s.cachet + SUM(spe.price) + v.price
      FROM shows AS s, events AS e, events_service_providers_serve AS spe, venues AS v
      WHERE e.show = s.id AND spe.event = e.id AND e.venue = v.id AND e.id =  t_e.id
    ) expenses
    FROM seats AS s, sectors_events_cost AS sec, tickets AS t
    WHERE s.sector = sep.sector AND t.seat = s.id
      AND sep.event = t_e.id AND t.event = t_e.id
  ) earnings 
  FROM events AS t_e 
  WHERE t_e.show = t_sh.id
    AND t_e.starts_at >= ...
    AND t_e.ends_at <= ...
) show_earnings
FROM show as t_sh
ORDER BY show_earnings DESC
LIMIT 1
