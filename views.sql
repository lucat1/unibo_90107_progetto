CREATE VIEW spesa_evento AS
SELECT
  e.id as evento,
  s.prezzo_siae + s.cachet + l.prezzo + SUM(efs.prezzo) AS "spesa"
FROM
  spettacolo AS s,
  evento AS e,
  evento_fornitore_servizio AS efs,
  luogo AS l
WHERE
  e.spettacolo = s.id
  AND efs.evento = e.id
  AND e.luogo = l.id
GROUP BY
  e.id,
  s.prezzo_siae,
  s.cachet,
  l.prezzo;

CREATE VIEW incasso_evento AS
SELECT
  b.evento,
  SUM(sec.prezzo) as "incasso"
FROM
  posto AS p,
  settore_evento_costo AS sec,
  biglietto AS b
WHERE
  p.settore = sec.settore
  AND b.posto = p.id
  AND sec.evento = b.evento
GROUP BY
  b.evento;

CREATE VIEW guadagno_evento AS
SELECT
  ie.evento evento,
  (ie.incasso - se.spesa) guadagno
FROM
  spesa_evento AS se NATURAL
  JOIN incasso_evento as ie
GROUP BY
  ie.evento,
  guadagno;

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
  s.id;

CREATE VIEW percentuale_biglietti_evento AS
SELECT
  be.evento AS evento,
  be.biglietti * 100.0 / SUM(s.capienza) AS percentuale
FROM
  settore_evento_costo AS sec,
  (
    SELECT
      b.evento AS evento,
      COUNT(*) AS biglietti
    FROM
      biglietto AS b
    GROUP BY
      b.evento
  ) AS be,
  settore AS s
WHERE
  s.id = sec.settore
  AND be.evento = sec.evento
GROUP BY
  be.evento,
  be.biglietti;