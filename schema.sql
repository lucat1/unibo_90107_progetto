CREATE TABLE artists (
  id serial,
  name text NOT NULL,
  
  PRIMARY KEY(id)
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

-- TODO
CREATE TYPE service_type AS ENUM ('stadium', 'theatre', 'arena');

CREATE TABLE venues (
  id serial,
  type venue_type NOT NULL,
  name text NOT NULL,
  coords point NOT NULL,
  price float NOT NULL,                                    -- per day

  CHECK(price >= 0),
  PRIMARY KEY (id),
  UNIQUE(name, coords)
);

CREATE TABLE events (
  id serial,
  show int NOT NULL,
  venue int NOT NULL,
  title text NOT NULL,
  starts_at timestamp NOT NULL,
  ends_at timestamp NOT NULL,

  PRIMARY KEY (id),
  FOREIGN KEY (show) REFERENCES shows,
  FOREIGN KEY (venue) REFERENCES venues
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
  row varchar(1) NOT NULL,
  col int NOT NULL,

  PRIMARY KEY (id),
  FOREIGN KEY (sector) REFERENCES venue_sectors
);

CREATE TABLE venue_sector_events_prices (
  sector int NOT NULL,
  event int NOT NULL,
  price float NOT NULL,

  CHECK(price > 0),
  UNIQUE (sector, event),
  FOREIGN KEY (sector) REFERENCES venue_sectors,
  FOREIGN KEY (event) REFERENCES events
);

CREATE TABLE service_providers (
  name varchar(256) NOT NULL,
  description text NOT NULL,

  PRIMARY KEY (name),
);

-- TODO
CREATE TYPE service_type AS ENUM ('audio', 'visual', 'guarding');

CREATE TABLE service_providers_events (
  provider varchar(256) NOT NULL,
  type service_type,
  event int NOT NULL,
  price float NOT NULL,                                    -- per hour

  CHECK(price > 0),
  FOREIGN KEY (provider) REFERENCES service_providers,
  FOREIGN KEY (event) REFERENCES events,
  UNIQUE (provider, type, event)
);

CREATE TABLE tickets (
  id serial,
  seat int NOT NULL,
  event int NOT NULL,

  PRIMARY KEY (id),
  FOREIGN KEY (seat) REFERENCES venue_sector_seats,
  FOREIGN KEY (event) REFERENCES events
);
