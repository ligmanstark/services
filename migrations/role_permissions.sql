CREATE TABLE role_permissions (
    role_id INT NOT NULL,
    permissions_id INT NOT NULL,
    PRIMARY KEY (role_id, permissions_id),
    FOREIGN KEY (role_id) REFERENCES role(id) ON DELETE CASCADE,
    FOREIGN KEY (permissions_id) REFERENCES permissions(id) ON DELETE CASCADE
)