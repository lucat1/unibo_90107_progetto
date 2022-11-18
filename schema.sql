CREATE TABLE events (
  id int NOT NULL AUTO_INCREMENT,
  show int NOT NULL,
  title text NOT NULL,
  location int NOT NULL,
  start datetime NOT NULL,
  end datetime NOT NULL,

  PRIMARY KEY (id),
  FOREIGN KEY (show) REFERENCES (shows),
  FOREIGN KEY (location) REFERENCES (locations)
)

CREATE TABLE shows (
  id int NOT NULL AUTO_INCREMENT,
  title text NOT NULL,
  artist int NOT NULL,
  siae_price float NOT NULL,
  cachet float NOT NULL,

  CHECK(cachet > 0 AND siae_price > 0),
  PRIMARY KEY (id),
  FOREIGN KEY (artist) REFERENCES (artists)
)

CREATE TABLE artists (
  id int NOT NULL AUTO_INCREMENT,
  name text NOT NULL,
  
  PRIMARY KEY(id)
)

CREATE TABLE locations (
  id int NOT NULL AUTO_INCREMENT,
  type int NOT NULL,
  name text NOT NULL,
  coords point NOT NULL,
  price float NOT NULL,                                    -- per hour

  CHECK(price > 0),
  PRIMARY KEY (id),
  FOREIGN KEY (type) REFERENCES (location_types)
)

CREATE TABLE location_sectors (
  id int NOT NULL AUTO_INCREMENT,
  location int NOT NULL,
  name text NOT NULL,
  capacity int NOT NULL,

  CHECK(capacity > 0),
  PRIMARY KEY (id),
  FOREIGN KEY (location) REFERENCES (locations)
)

CREATE TABLE location_sector_events (
  sector int NOT NULL,
  event int NOT NULL,
  price float64 NOT NULL,

  CHECK(price > 0),
  UNIQUE (sector, event),
  FOREIGN KEY (sector) REFERENCES (sectors),
  FOREIGN KEY (event) REFERENCES (events)
)

CREATE TABLE location_types (
  id int NOT NULL AUTO_INCREMENT,
  name varchar(256) NOT NULL,

  PRIMARY KEY (id)
)

CREATE TABLE service_providers (
  id int NOT NULL AUTO_INCREMENT,
  type int NOT NULL,
  price float NOT NULL,                                    -- per hour

  CHECK(price > 0),
  PRIMARY KEY (id),
  FOREIGN KEY (type) REFERENCES (service_types)
)

CREATE TABLE service_types (
  id int NOT NULL AUTO_INCREMENT,
  name varchar(256) NOT NULL,
  description text NOT NULL,

  PRIMARY KEY (id)
)

CREATE TABLE service_providers_events (
  provider int NOT NULL, 
  event int NOT NULL,

  FOREING KEY (provider) REFENCES (service_providers),
  FOREING KEY (event) REFENCES (events),
  UNIQUE (provider, event)
)

CREATE TABLE tickets (
  id int NOT NULL AUTO_INCREMENT,
  sector int NOT NULL,
  event int NOT NULL,

  PRIMARY KEY (id),
  FOREIGN KEY (sector) REFERENCES (location_sectors),
  FOREIGN KEY (event) REFERENCES (events)
)
