-- tested
SELECT DISTINCT
    e.titolo,
    l.nome
FROM
    evento AS e,
    luogo AS l,
    spettacolo AS s,
    artista AS a
WHERE
    l.id = e.luogo
    AND s.id = e.spettacolo
    AND s.artista = ...
