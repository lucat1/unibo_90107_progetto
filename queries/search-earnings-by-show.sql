SELECT
  SUM (
    SELECT
      SUM(sec.prezzo) - (
        SELECT
          s.prezzo_siae + s.cachet + SUM(efs.prezzo) + l.prezzo
        FROM
          spettacolo AS s,
          evento AS e,
          evento_fornitore_servizio AS efs,
          luogo AS l
        WHERE
          e.spettacolo = s.id
          AND efs.evento = e.id
          AND e.luogo = l.id
          AND e.id = t_e.id
      ) spese
    FROM
      posto AS p,
      settore_evento_costo AS sec,
      biglietto AS b
    WHERE
      p.sector = sec.sector
      AND b.seat = p.id
      AND sec.event = t_e.id
      AND b.event = t_e.id
  )
FROM
  evento AS t_e
WHERE
  t_e.spettacolo = ...