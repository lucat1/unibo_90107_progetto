SELECT
    t_l.nome,
    (
        SELECT
            SUM(capienza)
        FROM
            settore
        WHERE
            luogo = t_l.id
    ) capienza_totale
FROM
    luogo AS t_l
WHERE
    capienza_totale >= ...