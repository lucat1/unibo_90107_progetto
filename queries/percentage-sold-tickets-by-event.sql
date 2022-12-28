SELECT
  (
    SELECT
      COUNT(*)
    FROM
      biglietto
    WHERE
      evento = ...
  ) * 100.0 / SUM(s.capienza)
FROM
  settore_evento_costo AS sec,
  settore AS s
WHERE
  s.id = sec.settore
  AND sec.evento = ...