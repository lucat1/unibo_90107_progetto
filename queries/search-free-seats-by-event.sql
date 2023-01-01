-- tested
SELECT
  p.settore,
  p.fila,
  p.numero
FROM
  posto AS p,
  settore_evento_costo AS sec
WHERE
  p.settore = sec.settore
  AND p.id NOT IN (
    SELECT
      posto
    FROM
      biglietto
    WHERE
      evento = sec.evento
  )
  AND sec.evento = ...