INSERT INTO room (building_name, floor, name) VALUES
  ('片柳研究所棟', 11, 'テストルーム1'),

  ('講義棟', 2, '講義棟201'),
  ('講義棟', 2, '講義棟202'),
  ('講義棟', 2, '講義棟203'),
  ('講義棟', 2, '講義棟204'),

  ('講義棟', 3, '講義棟301'),
  ('講義棟', 3, '講義棟302'),
  ('講義棟', 3, '講義棟303'),
  ('講義棟', 3, '講義棟304');

INSERT INTO thing (room_id, thing_name) VALUES
  (1, 'Room1_yuuki'),
  (1, 'Room2_yuuki');
