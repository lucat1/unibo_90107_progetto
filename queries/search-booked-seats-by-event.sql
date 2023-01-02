SELECT
    p.settore,
    p.fila,
    p.numero
FROM
    posto AS p,
    settore_evento_costo AS sec,
    biglietto AS b
WHERE
    p.settore = sec.settore
    AND b.posto = p.id
    AND b.evento = sec.evento
    AND sec.evento = ...