export interface User {
  id: string;
  email: string;
  role: 'admin' | 'analyst';
  createdAt: string;
}

export interface IOC {
  id: string;
  type: 'ip' | 'domain' | 'url' | 'hash_md5' | 'hash_sha256' | 'email';
  value: string;
  tlpLevel: 'white' | 'green' | 'amber' | 'red';
  confidence: number;
  tags: string[];
  source: string;
  description: string;
  firstSeen: string | null;
  lastSeen: string | null;
  createdAt: string;
  updatedAt: string;
}

export interface ThreatActor {
  id: string;
  name: string;
  aliases: string[];
  sophistication: 'minimal' | 'intermediate' | 'advanced' | 'expert';
  resourceLevel: string;
  primaryMotivation: string;
  countryCode: string;
  description: string;
  campaignCount: number;
  createdAt: string;
  updatedAt: string;
}

export interface Campaign {
  id: string;
  name: string;
  description: string;
  firstSeen: string | null;
  lastSeen: string | null;
  objective: string;
  threatActorId: string | null;
  threatActorName: string | null;
  threatActorSophistication: string | null;
  createdAt: string;
  updatedAt: string;
}

export interface ATTACKMapping {
  id: string;
  techniqueId: string;
  techniqueName: string;
  tactic: string;
  platform: string[];
  entityType: 'ioc' | 'threat_actor' | 'campaign';
  entityId: string;
  createdAt: string;
}

export interface STIXBundle {
  id: string;
  specVersion: string;
  bundleJson: any;
  createdAt: string;
}

export interface PaginatedResponse<T> {
  data: T[];
  total: number;
  page: number;
  limit: number;
}

export interface ConfidenceBucket {
  bucket: string;
  count: number;
}

export interface TypeBreakdownItem {
  type: string;
  count: number;
}

export interface DashboardStats {
  totalIOCs: number;
  totalThreatActors: number;
  totalCampaigns: number;
  totalAttackMappings: number;
  confidenceDistribution: ConfidenceBucket[];
  typeBreakdown: TypeBreakdownItem[];
  recentIOCs: IOC[];
}
