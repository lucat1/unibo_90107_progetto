SELECT
  t_sp.titolo,
  (
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
          p.settore = sec.settore
          AND b.posto = p.id
          AND sec.evento = t_e.id
          AND b.evento = t_e.id
      ) guadagni
    FROM
      evento AS t_e
    WHERE
      t_e.spettacolo = t_sp.id
      AND t_e.inizio >=...
      AND t_e.fine <=...
  ) guadagni_spettacolo
FROM
  spettacolo as t_sp
ORDER BY
  guadagni_spettacolo DESC
LIMIT
  1