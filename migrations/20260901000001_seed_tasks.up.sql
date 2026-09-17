INSERT INTO tasks (id, title, completed, created_at, updated_at) VALUES
  ('095bfb69-202a-4b25-a23b-4d65c6f0b3c7', 'Preparar la reunión del lunes', false, '2026-09-01 09:00:00', '2026-09-01 09:00:00'),
  ('3fa0d1cd-4b3e-4b1f-b577-91a9903ebb5f', 'Enviar el informe mensual', false, '2026-09-01 09:05:00', '2026-09-01 09:05:00'),
  ('e172e765-085c-47c9-9825-62da6f220027', 'Renovar el dominio de la web', true, '2026-09-01 09:10:00', '2026-09-01 09:10:00')
ON CONFLICT (id) DO NOTHING;
