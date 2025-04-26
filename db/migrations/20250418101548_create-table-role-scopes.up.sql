CREATE TABLE IF NOT EXISTS role_scopes (
    role_scope_id SERIAL PRIMARY KEY,
    role_id int NOT NULL,
    scope_id int NOT NULL,
    CONSTRAINT fk_scope
      FOREIGN KEY(scope_id)
        REFERENCES scopes(id)
);