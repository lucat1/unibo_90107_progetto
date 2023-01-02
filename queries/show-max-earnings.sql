-- tested and exported
SELECT
  gs.spettacolo spettacolo,
  gs.guadagno guadagno
FROM
  guadagno_spettacolo as gs
ORDER BY
  guadagno DESC
LIMIT
  1