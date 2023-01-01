-- tested
SELECT
  be.biglietti * 100.0 / SUM(s.capienza)
FROM
  settore_evento_costo AS sec,
  (
    SELECT
      b.evento AS evento,
      COUNT(*) AS biglietti
    FROM
      biglietto AS b
    GROUP BY
      b.evento
  ) AS be,
  settore AS s
WHERE
  s.id = sec.settore
  AND be.evento = sec.evento
  AND sec.evento = ...
GROUP BY
  be.biglietti