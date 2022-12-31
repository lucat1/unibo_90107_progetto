-- tested
SELECT
  e.titolo
FROM
  evento AS e
WHERE
  (
    SELECT
      COUNT(p.id)
    FROM
      posto AS p,
      settore_evento_costo AS sec
    WHERE
      p.settore = sec.settore
      AND p.id NOT IN (
        SELECT
          posto
        FROM
          biglietto
        WHERE
          evento = e.id
      )
      AND sec.evento = e.id
  ) = 0
