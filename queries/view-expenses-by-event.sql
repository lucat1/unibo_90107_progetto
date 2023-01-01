-- tested
CREATE VIEW spesa_evento AS
SELECT
    e.id as evento,
    s.prezzo_siae + s.cachet + l.prezzo + SUM(efs.prezzo) AS "spesa"
FROM
    spettacolo AS s,
    evento AS e,
    evento_fornitore_servizio AS efs,
    luogo AS l
WHERE
    e.spettacolo = s.id
    AND efs.evento = e.id
    AND e.luogo = l.id
GROUP BY
    e.id,
    s.prezzo_siae,
    s.cachet,
    l.prezzo