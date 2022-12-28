SELECT
    p.fila,
    p.numero
FROM
    posto AS p,
    settore_evento_costo AS sec,
    biglietto AS b
WHERE
    p.sector = sec.sector
    AND b.seat = p.id
    AND b.event = sec.event
    AND sec.event = ...