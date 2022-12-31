-- tested
CREATE view incasso_evento AS
  SELECT
    b.evento, SUM(sec.prezzo) as "incasso"
  FROM
    posto AS p,
    settore_evento_costo AS sec,
    biglietto AS b
  WHERE
    p.settore = sec.settore
    AND b.posto = p.id
    AND sec.evento = b.evento
  GROUP BY b.evento
