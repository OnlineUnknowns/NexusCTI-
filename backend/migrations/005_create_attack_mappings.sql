CREATE TABLE IF NOT EXISTS attack_mappings (
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  technique_id TEXT NOT NULL,
  technique_name TEXT NOT NULL,
  tactic TEXT NOT NULL,
  platform TEXT[] DEFAULT '{}',
  entity_type TEXT NOT NULL CHECK (entity_type IN ('ioc','threat_actor','campaign')),
  entity_id UUID NOT NULL,
  created_at TIMESTAMPTZ DEFAULT NOW()
);
