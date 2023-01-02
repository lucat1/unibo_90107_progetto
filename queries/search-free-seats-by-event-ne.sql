SELECT
  p.settore,
  p.fila,
  p.numero
FROM
  posto AS p,
  settore_evento_costo AS sec
WHERE
  p.settore = sec.settore
  AND NOT EXISTS (
    SELECT
      posto
    FROM
      biglietto
    WHERE
      posto = p.id
      AND evento = sec.evento
  )
  AND sec.evento = ...
