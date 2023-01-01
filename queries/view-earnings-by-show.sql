-- tested
CREATE VIEW guadagno_spettacolo AS
SELECT
  s.id spettacolo,
  SUM(ge.guadagno) guadagno
FROM
  spettacolo AS s,
  evento AS e,
  guadagno_evento AS ge
WHERE
  e.id = ge.evento
  AND e.spettacolo = s.id
GROUP BY
  s.id