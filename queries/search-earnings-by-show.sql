-- tested 
SELECT
  SUM(ge.guadagno) guadagno
FROM
  spettacolo AS s,
  evento AS e,
  guadagno_evento AS ge
WHERE
  e.id = ge.evento
  AND e.spettacolo = s.id
  AND s.id = ...