CREATE TABLE IF NOT EXISTS role_scopes (
    role_scope_id SERIAL PRIMARY KEY,
    role_id int NOT NULL,
    scope_id int NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE,
    CONSTRAINT fk_scope
      FOREIGN KEY(scope_id)
        REFERENCES scopes(id)
);