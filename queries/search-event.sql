SELECT
   e.titolo, e.inizio, e.fine, l.nome
FROM
  evento AS e,
  luogo AS l
WHERE
  e.luogo = l.id AND
  e.id = ...
