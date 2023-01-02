SELECT
  (ie.incasso - se.spesa) guadagno
FROM
  spesa_evento AS se NATURAL
  JOIN incasso_evento as ie
WHERE
  se.evento = ...
