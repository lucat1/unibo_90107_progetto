CREATE TABLE events (
  id serial,
  show int NOT NULL,
  title text NOT NULL,
  venue int NOT NULL,
  start datetime NOT NULL,
  end datetime NOT NULL,

  PRIMARY KEY (id),
  FOREIGN KEY (show) REFERENCES shows,
  FOREIGN KEY (venue) REFERENCES venues
);

CREATE TABLE shows (
  id serial,
  title text NOT NULL,
  artist int NOT NULL,
  siae_price float NOT NULL,
  cachet float NOT NULL,

  CHECK(cachet > 0 AND siae_price > 0),
  PRIMARY KEY (id),
  FOREIGN KEY (artist) REFERENCES artists
);

CREATE TABLE artists (
  id serial,
  name text NOT NULL,
  
  PRIMARY KEY(id)
);

CREATE TABLE venues (
  id serial,
  type int NOT NULL,
  name text NOT NULL,
  coords point NOT NULL,
  price float NOT NULL,                                    -- per hour

  CHECK(price > 0),
  PRIMARY KEY (id),
  FOREIGN KEY (type) REFERENCES venue_types
);

CREATE TABLE venue_sectors (
  id serial,
  venue int NOT NULL,
  name text NOT NULL,
  capacity int NOT NULL,

  CHECK(capacity > 0),
  PRIMARY KEY (id),
  FOREIGN KEY (venue) REFERENCES venues
);

CREATE TABLE venue_sector_seats (
  id serial,
  sector int NOT NULL,
  row char NOT NULL,
  column int NOT NULL,

  PRIMARY KEY (id),
  FOREIGN KEY (sector) REFERENCES venue_sectors
)

CREATE TABLE venue_sector_events (
  sector int NOT NULL,
  event int NOT NULL,
  price float64 NOT NULL,

  CHECK(price > 0),
  UNIQUE (sector, event),
  FOREIGN KEY (sector) REFERENCES sectors,
  FOREIGN KEY (event) REFERENCES events
);

CREATE TABLE venue_types (
  id serial,
  name varchar(256) NOT NULL,

  PRIMARY KEY (id)
);

CREATE TABLE service_providers (
  id serial,
  type int NOT NULL,
  price float NOT NULL,                                    -- per hour

  CHECK(price > 0),
  PRIMARY KEY (id),
  FOREIGN KEY (type) REFERENCES service_types
);

CREATE TABLE service_types (
  id serial,
  name varchar(256) NOT NULL,
  description text NOT NULL,

  PRIMARY KEY (id)
);

CREATE TABLE service_providers_events (
  provider int NOT NULL, 
  event int NOT NULL,

  FOREIGN KEY (provider) REFERENCES service_providers,
  FOREIGN KEY (event) REFERENCES events,
  UNIQUE (provider, event)
);

CREATE TABLE tickets (
  id serial,
  seat int NOT NULL,
  event int NOT NULL,

  PRIMARY KEY (id),
  FOREIGN KEY (seat) REFERENCES venue_sector_seats,
  FOREIGN KEY (event) REFERENCES events
);
