-- tested
CREATE VIEW guadagno_evento AS
SELECT
  ie.evento evento,
  (ie.incasso - se.spesa) guadagno
FROM
  spesa_evento AS se NATURAL
  JOIN incasso_evento as ie
GROUP BY
  ie.evento,
  guadagno