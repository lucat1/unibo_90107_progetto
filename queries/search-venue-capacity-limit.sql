SELECT
    l.id,
    l.nome,
    SUM(s.capienza) capienza_totale
FROM
    settore AS s,
    luogo AS l
WHERE
    s.luogo = l.id
GROUP BY
    l.id
HAVING
    SUM(s.capienza) >= ...