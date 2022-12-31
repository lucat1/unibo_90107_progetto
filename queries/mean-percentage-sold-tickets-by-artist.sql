-- tested
SELECT
  AVG(
    pbe.percentuale
  )
FROM
  spettacolo AS s,
  evento AS e,
  percentuale_biglietti_evento AS pbe
WHERE
  s.id = e.spettacolo 
  AND pbe.evento = e.id
  AND s.artista = ...
