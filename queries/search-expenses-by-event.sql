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