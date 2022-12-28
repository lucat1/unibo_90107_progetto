SELECT
  AVG(
    (
      SELECT
        COUNT(*)
      FROM
        biglietto
      WHERE
        evento = e.id
    ) * 100.0 / SUM(s.capienza)
  )
FROM
  settore_evento_costo AS sec,
  settore AS s,
  spettacolo AS sp,
  evento AS e
WHERE
  s.id = sec.settore
  AND sec.evento = e.id
  AND sp.artista = ...
  AND sp.id = e.spettacolo