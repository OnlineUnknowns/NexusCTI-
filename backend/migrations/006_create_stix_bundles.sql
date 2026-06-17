CREATE TABLE IF NOT EXISTS stix_bundles (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  spec_version TEXT NOT NULL DEFAULT '2.1',
  bundle_json JSONB NOT NULL,
  created_at TIMESTAMPTZ DEFAULT NOW()
);
