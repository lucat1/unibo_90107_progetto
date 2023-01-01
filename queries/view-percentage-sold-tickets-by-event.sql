CREATE VIEW percentuale_biglietti_evento AS
SELECT
  be.evento AS evento,
  be.biglietti * 100.0 / SUM(s.capienza) AS percentuale
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
GROUP BY
  be.evento,
  be.biglietti