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
      AND e.id = ...
  ) spese
FROM
  posto AS p,
  settore_evento_costo AS sec,
  biglietto AS b
WHERE
  p.settore = sec.settore
  AND b.posto = p.id
  AND sec.evento = ...
  AND b.evento = ...