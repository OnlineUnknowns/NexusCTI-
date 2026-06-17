DO $$
BEGIN
  IF NOT EXISTS (SELECT 1 FROM threat_actors) THEN
    -- Insert Threat Actors
    INSERT INTO threat_actors (id, name, aliases, sophistication, resource_level, primary_motivation, country_code, description) VALUES
      ('8d89a42f-763b-486b-a81d-b6329a1b0281', 'APT28', ARRAY['Fancy Bear', 'Sofacy'], 'expert', 'government', 'espionage', 'RU', 'A Russian cyber espionage group active since at least 2004.'),
      ('cf9e782e-9d22-4aee-b248-cbbf102f9011', 'Lazarus Group', ARRAY['Hidden Cobra', 'Guardians of Peace'], 'advanced', 'government', 'financial-gain', 'KP', 'A state-sponsored cyber group originating from North Korea.'),
      ('d9426f8d-4e92-4a0b-9ef1-f09c6ee02821', 'Sandworm', ARRAY['TeleBots', 'Voodoo Bear'], 'expert', 'government', 'sabotage', 'RU', 'A Russian cyber warfare unit operating under GRU direction.');

    -- Insert Campaigns
    INSERT INTO campaigns (id, name, description, first_seen, last_seen, objective, threat_actor_id) VALUES
      ('12d23b7b-23f2-49bd-831e-128a1c905b22', 'Operation Ghost', 'Cyber espionage campaign targeting European government agencies.', '2023-01-15 00:00:00+00', '2023-08-30 00:00:00+00', 'Espionage', '8d89a42f-763b-486b-a81d-b6329a1b0281'),
      ('4466b09c-e35b-426b-aefc-34ea1c841bb2', 'DarkSeoul', 'Coordinated cyber attacks targeting South Korean banks and broadcasters.', '2022-03-01 00:00:00+00', '2022-06-15 00:00:00+00', 'Disruption and damage', 'cf9e782e-9d22-4aee-b248-cbbf102f9011');

    -- Insert IOCs
    INSERT INTO iocs (id, type, value, tlp_level, confidence, tags, source, description, first_seen, last_seen) VALUES
      ('01f11a43-e69d-429f-a89c-d0961de3331b', 'ip', '185.220.101.45', 'red', 90, ARRAY['tor-exit', 'malicious'], 'US-CERT', 'Tor exit node used in APT28 operations.', '2023-02-10 12:00:00+00', '2023-08-15 14:30:00+00'),
      ('5c2020e9-b5fe-4b14-8742-5dbbf48301cb', 'domain', 'evil-domain.ru', 'amber', 80, ARRAY['c2', 'phishing'], 'FireEye', 'Phishing domain associated with credential harvesting.', '2023-03-05 09:00:00+00', '2023-08-20 18:00:00+00'),
      ('e5f992aa-cfbc-418d-9fe8-251de0023a1c', 'url', 'http://malware.biz/payload.exe', 'red', 95, ARRAY['malware-delivery'], 'CrowdStrike', 'Malicious payload delivery URL.', '2022-04-12 11:22:00+00', '2022-06-01 10:15:00+00'),
      ('d62e78fa-a10c-4467-bc22-38d21c9a1ffc', 'hash_md5', '44d88612fea8a8f36de82e1278abb02f', 'amber', 70, ARRAY['trojan'], 'Mandiant', 'MD5 hash of a custom backdoor payload.', '2022-03-15 08:00:00+00', '2022-05-10 09:00:00+00'),
      ('2a04871e-08bc-4674-bf4b-e60d195c1bcf', 'hash_sha256', 'a591a6d40bf420404a011733cfb7b190d62e78faa10c4467bc2238d21c9a1ffc', 'green', 60, ARRAY['downloader'], 'Kaspersky', 'SHA256 of the loader executable.', '2023-06-01 00:00:00+00', '2023-08-25 23:59:00+00');

    -- Insert ATT&CK Mappings
    INSERT INTO attack_mappings (technique_id, technique_name, tactic, platform, entity_type, entity_id) VALUES
      ('T1059', 'Command and Scripting Interpreter', 'Execution', ARRAY['Windows', 'Linux'], 'threat_actor', '8d89a42f-763b-486b-a81d-b6329a1b0281'),
      ('T1078', 'Valid Accounts', 'Defense Evasion', ARRAY['Windows', 'Active Directory'], 'threat_actor', 'cf9e782e-9d22-4aee-b248-cbbf102f9011'),
      ('T1566', 'Phishing', 'Initial Access', ARRAY['Windows', 'macOS', 'Office 365'], 'threat_actor', 'd9426f8d-4e92-4a0b-9ef1-f09c6ee02821'),
      ('T1190', 'Exploit Public-Facing Application', 'Initial Access', ARRAY['Linux', 'Network'], 'campaign', '12d23b7b-23f2-49bd-831e-128a1c905b22'),
      ('T1027', 'Obfuscated Files or Information', 'Defense Evasion', ARRAY['Windows', 'macOS'], 'campaign', '4466b09c-e35b-426b-aefc-34ea1c841bb2');
  END IF;
END $$;
