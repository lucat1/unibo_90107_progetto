-- tested
SELECT
  SUM(sec.prezzo) as "incasso totale"
FROM
  posto AS p,
  settore_evento_costo AS sec,
  biglietto AS b
WHERE
  p.settore = sec.settore
  AND b.posto = p.id
  AND sec.evento = b.evento
  AND b.evento = ...
