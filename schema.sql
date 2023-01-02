CREATE TABLE persona (
  id serial,
  nome text NOT NULL,
  cognome text NOT NULL,
  data_nascita date NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE gruppo (
  id serial,
  data_formazione date NOT NULL,
  PRIMARY KEY (id)
);

CREATE TABLE persona_gruppo_appartenenza (
  persona int NOT NULL,
  gruppo int NOT NULL,
  FOREIGN KEY (persona) REFERENCES persona,
  FOREIGN KEY (gruppo) REFERENCES gruppo,
  UNIQUE (persona, gruppo)
);

CREATE TABLE artista (
  id serial,
  nome_arte text NOT NULL,
  persona int,
  gruppo int,
  CHECK(
    (
      persona IS NOT NULL
      AND gruppo is NULL
    )
    OR (
      gruppo IS NOT NULL
      AND persona is NULL
    )
  ),
  PRIMARY KEY (id),
  FOREIGN KEY (persona) REFERENCES persona,
  FOREIGN KEY (gruppo) REFERENCES gruppo,
  UNIQUE (nome_arte)
);

CREATE TABLE spettacolo (
  id serial,
  titolo text NOT NULL,
  artista int NOT NULL,
  prezzo_siae money NOT NULL,
  cachet money NOT NULL,
  CHECK (
    cachet >= CAST(0 AS money)
    AND prezzo_siae >= CAST(0 AS money)
  ),
  PRIMARY KEY (id),
  FOREIGN KEY (artista) REFERENCES artista,
  UNIQUE (titolo, artista)
);

CREATE TYPE tipo_luogo AS ENUM (
  'arena',
  'palazzetto',
  'parco',
  'piazza',
  'stadio',
  'teatro'
);

CREATE TABLE luogo (
  id serial,
  tipo tipo_luogo NOT NULL,
  nome text NOT NULL,
  indirizzo text NOT NULL,
  citta text NOT NULL,
  prezzo money NOT NULL,
  -- per day
  CHECK (prezzo >= CAST(0 AS money)),
  PRIMARY KEY (id),
  UNIQUE (nome, indirizzo, citta)
);

CREATE TABLE evento (
  id serial,
  spettacolo int NOT NULL,
  luogo int NOT NULL,
  titolo text NOT NULL,
  inizio timestamp NOT NULL,
  fine timestamp NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (spettacolo) REFERENCES spettacolo,
  FOREIGN KEY (luogo) REFERENCES luogo,
  UNIQUE (titolo, inizio, luogo),
  CHECK(fine > inizio)
);

CREATE TABLE settore (
  id serial,
  luogo int NOT NULL,
  nome text NOT NULL,
  capienza int NOT NULL,
  CHECK (capienza > 0),
  PRIMARY KEY (id),
  FOREIGN KEY (luogo) REFERENCES luogo,
  UNIQUE (luogo, nome)
);

CREATE TABLE posto (
  id serial,
  settore int NOT NULL,
  fila varchar(1) NOT NULL,
  numero int NOT NULL,
  PRIMARY KEY (id),
  FOREIGN KEY (settore) REFERENCES settore,
  UNIQUE (settore, fila, numero)
);

CREATE TABLE settore_evento_costo (
  settore int NOT NULL,
  evento int NOT NULL,
  -- prezzo per biglietto
  prezzo money NOT NULL,
  CHECK (prezzo >= CAST(0 AS money)),
  UNIQUE (settore, evento),
  FOREIGN KEY (settore) REFERENCES settore,
  FOREIGN KEY (evento) REFERENCES evento
);

CREATE TABLE fornitore (
  id serial,
  nome text NOT NULL,
  descrizione text NOT NULL,
  PRIMARY KEY (id),
  UNIQUE (nome)
);

CREATE TYPE tipo_servizio AS ENUM (
  'audio',
  'biglietteria',
  'luci',
  'maschere',
  'sicurezza',
  'regia'
);

CREATE TABLE evento_fornitore_servizio (
  fornitore int NOT NULL,
  tipo tipo_servizio,
  evento int NOT NULL,
  prezzo money NOT NULL,
  -- for the whole job
  CHECK (prezzo > CAST(0 AS money)),
  FOREIGN KEY (fornitore) REFERENCES fornitore,
  FOREIGN KEY (evento) REFERENCES evento,
  UNIQUE (fornitore, tipo, evento)
);

CREATE TABLE biglietto (
  codice serial,
  nominativo text NOT NULL,
  posto int NOT NULL,
  evento int NOT NULL,
  PRIMARY KEY (codice),
  FOREIGN KEY (posto) REFERENCES posto,
  FOREIGN KEY (evento) REFERENCES evento
);