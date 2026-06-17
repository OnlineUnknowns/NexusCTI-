CREATE TABLE IF NOT EXISTS iocs (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  type TEXT NOT NULL CHECK (type IN ('ip','domain','url','hash_md5','hash_sha256','email')),
  value TEXT NOT NULL,
  tlp_level TEXT NOT NULL DEFAULT 'white' CHECK (tlp_level IN ('white','green','amber','red')),
  confidence INTEGER NOT NULL DEFAULT 50 CHECK (confidence >= 0 AND confidence <= 100),
  tags TEXT[] DEFAULT '{}',
  source TEXT,
  description TEXT,
  first_seen TIMESTAMPTZ,
  last_seen TIMESTAMPTZ,
  search_vector tsvector GENERATED ALWAYS AS (
    to_tsvector('english', coalesce(value,'') || ' ' || coalesce(description,'') || ' ' || coalesce(source,''))
  ) STORED,
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW()
);
CREATE INDEX IF NOT EXISTS iocs_search_idx ON iocs USING GIN(search_vector);
CREATE INDEX IF NOT EXISTS iocs_type_idx ON iocs(type);
CREATE INDEX IF NOT EXISTS iocs_tlp_idx ON iocs(tlp_level);
