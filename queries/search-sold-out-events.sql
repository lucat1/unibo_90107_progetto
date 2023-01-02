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
      AND sec.evento = e.id
      AND NOT EXISTS (
        SELECT
          posto
        FROM
          biglietto
        WHERE
          evento = e.id
          AND posto = p.id
      )
  ) = 0
