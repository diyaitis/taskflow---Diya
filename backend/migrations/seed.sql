INSERT INTO users VALUES (
  gen_random_uuid(),
  'Test User',
  'test@example.com',
  '$2a$12$hashedpassword',
  NOW()
);