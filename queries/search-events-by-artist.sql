SELECT
    DISTINCT e.titolo ,
    l.nome luogo
FROM
    evento AS e,
    luogo AS l,
    spettacolo AS s
WHERE
    l.id = e.luogo
    AND s.id = e.spettacolo
    AND s.artista = ...
