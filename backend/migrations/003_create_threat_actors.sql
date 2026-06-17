CREATE TABLE IF NOT EXISTS threat_actors (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  name TEXT NOT NULL,
  aliases TEXT[] DEFAULT '{}',
  sophistication TEXT CHECK (sophistication IN ('minimal','intermediate','advanced','expert')),
  resource_level TEXT,
  primary_motivation TEXT,
  country_code TEXT,
  description TEXT,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW()
);
