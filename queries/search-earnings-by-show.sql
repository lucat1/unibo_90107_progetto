SELECT
  SUM(ge.guadagno) guadagno
FROM
  evento AS e,
  guadagno_evento AS ge
WHERE
  e.id = ge.evento
  AND e.spettacolo = ...
