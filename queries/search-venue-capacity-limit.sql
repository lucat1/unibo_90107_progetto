SELECT
    nome,
    SUM(capienza) capienza_totale
FROM
    settore
GROUP BY
    luogo
WHERE
    capienza_totale >= ...;