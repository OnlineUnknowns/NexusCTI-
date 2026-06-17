<div align="center">

<!-- Animated Banner -->
<img src="https://capsule-render.vercel.app/api?type=venom&color=0:0d1117,50:1a1f2e,100:f78166&height=280&section=header&text=OpenCTI%20Lite&fontSize=80&fontColor=f78166&animation=twinkling&fontAlignY=55&desc=Enterprise%20Threat%20Intelligence%20Platform&descAlignY=75&descSize=22&descColor=8b949e" width="100%"/>

<!-- Animated typing -->
<a href="https://git.io/typing-svg"><img src="https://readme-typing-svg.demolab.com?font=JetBrains+Mono&weight=700&size=22&duration=3000&pause=1000&color=F78166&center=true&vCenter=true&multiline=true&repeat=true&width=700&height=100&lines=IOC+Management+%E2%80%A2+Threat+Actors+%E2%80%A2+Campaigns;STIX+2.1+%2F+TAXII+%E2%80%A2+MITRE+ATT%26CK+Mapping;Enterprise+Security+Intelligence+at+Scale" alt="Typing SVG" /></a>

<br/>

<!-- Badges Row 1 -->
<p>
  <img src="https://img.shields.io/badge/Go-1.21+-00ADD8?style=for-the-badge&logo=go&logoColor=white"/>
  <img src="https://img.shields.io/badge/React-18-61DAFB?style=for-the-badge&logo=react&logoColor=black"/>
  <img src="https://img.shields.io/badge/PostgreSQL-15-4169E1?style=for-the-badge&logo=postgresql&logoColor=white"/>
  <img src="https://img.shields.io/badge/Redis-7-DC382D?style=for-the-badge&logo=redis&logoColor=white"/>
  <img src="https://img.shields.io/badge/Docker-Compose-2496ED?style=for-the-badge&logo=docker&logoColor=white"/>
</p>

<!-- Badges Row 2 -->
<p>
  <img src="https://img.shields.io/badge/STIX-2.1-f78166?style=for-the-badge&logo=databricks&logoColor=white"/>
  <img src="https://img.shields.io/badge/TAXII-2.1-f78166?style=for-the-badge&logo=databricks&logoColor=white"/>
  <img src="https://img.shields.io/badge/MITRE%20ATT%26CK-Mapped-red?style=for-the-badge&logo=target&logoColor=white"/>
  <img src="https://img.shields.io/badge/JWT-Auth-000000?style=for-the-badge&logo=jsonwebtokens&logoColor=white"/>
  <img src="https://img.shields.io/badge/TypeScript-Strict-3178C6?style=for-the-badge&logo=typescript&logoColor=white"/>
</p>

<!-- Stats -->
<p>
  <img src="https://img.shields.io/github/stars/your-org/opencti-lite?style=for-the-badge&color=f78166&labelColor=0d1117"/>
  <img src="https://img.shields.io/github/forks/your-org/opencti-lite?style=for-the-badge&color=8b949e&labelColor=0d1117"/>
  <img src="https://img.shields.io/github/issues/your-org/opencti-lite?style=for-the-badge&color=3fb950&labelColor=0d1117"/>
  <img src="https://img.shields.io/github/license/your-org/opencti-lite?style=for-the-badge&color=d29922&labelColor=0d1117"/>
</p>

<br/>

<!-- Quick links -->
[**📖 Documentation**](#-documentation) • [**🚀 Quick Start**](#-quick-start) • [**🏗️ Architecture**](#️-architecture) • [**🛡️ Features**](#️-features) • [**🤝 Contributing**](#-contributing)

</div>

---

<br/>

<!-- Hero visual -->
<div align="center">

```
┌─────────────────────────────────────────────────────────────────────┐
│                                                                     │
│   🛡️  OpenCTI Lite  —  Real-Time Cyber Threat Intelligence         │
│                                                                     │
│   ┌──────────┐  ┌──────────┐  ┌──────────┐  ┌──────────────────┐  │
│   │  IOC     │  │ Threat   │  │Campaign  │  │  ATT&CK Matrix   │  │
│   │ Engine   │→ │ Actors   │→ │Tracking  │→ │  T1059 • T1078  │  │
│   │ 6 Types  │  │ 3 Levels │  │Timeline  │  │  T1566 • T1190  │  │
│   └──────────┘  └──────────┘  └──────────┘  └──────────────────┘  │
│         │              │              │                │            │
│         └──────────────┴──────────────┴────────────────┘           │
│                              ↓                                      │
│                    ┌─────────────────┐                              │
│                    │  STIX 2.1 Bundle│                              │
│                    │  TAXII 2.1 Feed │                              │
│                    └─────────────────┘                              │
│                                                                     │
└─────────────────────────────────────────────────────────────────────┘
```

</div>

<br/>

---

## 🎯 What is OpenCTI Lite?

**OpenCTI Lite** is a production-grade, open-source **Cyber Threat Intelligence (CTI) platform** built for security operations teams, threat researchers, and enterprise SOCs. It delivers the core intelligence workflows of platforms costing **$500K+/year** — packaged as a self-hosted, open-source solution.

<table>
<tr>
<td width="50%">

### 🔴 The Problem
- Commercial CTI platforms cost **$200K–$2M/year**
- Open-source alternatives are complex to deploy
- No single platform covers IOCs + actors + campaigns + ATT&CK + STIX natively
- Existing tools lack developer-friendly APIs

</td>
<td width="50%">

### 🟢 The Solution
- **Free & self-hosted** — own your intelligence data
- **5-minute deployment** via Docker Compose
- Full **STIX 2.1 / TAXII 2.1** compliance
- **MITRE ATT&CK** matrix integration out of the box
- Clean REST API — integrate with any SIEM/SOAR

</td>
</tr>
</table>

---

## 🛡️ Features

<div align="center">

### Core Intelligence Modules

</div>

<table>
<tr>
<td align="center" width="25%">

### 🔍
### IOC Management

Full lifecycle management for **6 indicator types**

`IP` `Domain` `URL` `MD5` `SHA256` `Email`

TLP classification, confidence scoring, full-text search, bulk CSV import

</td>
<td align="center" width="25%">

### 👤
### Threat Actors

Profile nation-state and criminal groups with **sophistication ratings**, country attribution, aliases, and motivation tracking

`APT28` `Lazarus` `Sandworm` and beyond

</td>
<td align="center" width="25%">

### 🎯
### Campaigns

Timeline-based campaign tracking linked to threat actors with objectives, date ranges, and associated TTPs

Visual timeline UI with actor attribution

</td>
<td align="center" width="25%">

### 🗺️
### ATT&CK Mapping

Interactive MITRE ATT&CK matrix overlaid with your real intelligence data — map techniques to IOCs, actors, and campaigns

14 tactics · 200+ techniques

</td>
</tr>
</table>

<br/>

<table>
<tr>
<td align="center" width="50%">

### 📦 STIX 2.1 / TAXII 2.1

```json
{
  "type": "bundle",
  "spec_version": "2.1",
  "objects": [
    { "type": "indicator", "pattern": "[ipv4-addr:value = '185.220.101.45']" },
    { "type": "threat-actor", "name": "APT28" },
    { "type": "campaign", "name": "Operation Ghost" },
    { "type": "relationship", "relationship_type": "attributed-to" }
  ]
}
```

Export your entire threat library as a STIX bundle. Share via TAXII feed. Import intelligence from external sources in one click.

</td>
<td align="center" width="50%">

### ⚡ Performance at Scale

```
Benchmark (10,000 IOCs)
─────────────────────────────────────
List  (cached)     →   < 5ms   ✅
Search (full-text) →  < 30ms   ✅
STIX Export        →  < 200ms  ✅
Bulk Import 1K IOC →  < 800ms  ✅
─────────────────────────────────────
Redis TTL Cache: 5 min
Rate Limiting:   100 req/min/IP
Full-text index: PostgreSQL tsvector
```

</td>
</tr>
</table>

---

## 🚀 Quick Start

### Prerequisites

```bash
# Required
docker --version    # Docker 24+
docker compose version  # Compose v2+
go version          # Go 1.21+
node --version      # Node 20+
```

### 1-Minute Deploy

```bash
# Clone the repository
git clone https://github.com/your-org/opencti-lite.git
cd opencti-lite

# Start infrastructure
docker compose up -d

# Start backend
cd backend
go mod tidy
go run main.go &

# Start frontend (new terminal)
cd ../frontend
npm install
npm run dev
```

<div align="center">

```
✅  PostgreSQL  →  localhost:5432
✅  Redis       →  localhost:6379
✅  Backend API →  http://localhost:8080
✅  Frontend UI →  http://localhost:5173
```

</div>

### First Login

```bash
# Register your admin account (first user = admin role)
curl -X POST http://localhost:8080/api/v1/auth/register \
  -H "Content-Type: application/json" \
  -d '{"email": "admin@yourorg.com", "password": "your-secure-password"}'

# Response:
# { "token": "eyJhbGci...", "user": { "role": "admin" } }
```

> **🌱 Seed Data Included** — The platform auto-seeds APT28, Lazarus Group, Sandworm, 5 IOCs, 2 campaigns, and ATT&CK mappings on first run.

---

## 🏗️ Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                         CLIENT LAYER                            │
│   React 18 + TypeScript + Vite + Tailwind + React Query        │
│   ┌────────┐ ┌──────────┐ ┌──────────┐ ┌────────┐ ┌────────┐  │
│   │Dashboard│ │   IOCs   │ │  Actors  │ │Campaigns│ │ STIX  │  │
│   └────────┘ └──────────┘ └──────────┘ └────────┘ └────────┘  │
└─────────────────────────┬───────────────────────────────────────┘
                          │  REST API + JWT
┌─────────────────────────▼───────────────────────────────────────┐
│                         API LAYER (Go/Gin)                      │
│  ┌──────────┐ ┌──────────┐ ┌──────────┐ ┌────────────────────┐ │
│  │   Auth   │ │Handlers  │ │Middleware│ │    STIX Service    │ │
│  │  (JWT)   │ │(CRUD+Bulk│ │(Rate Lim)│ │  Export / Import  │ │
│  └──────────┘ └──────────┘ └──────────┘ └────────────────────┘ │
│  ┌─────────────────┐  ┌──────────────────────────────────────┐  │
│  │  Cache Service  │  │         Search Service               │  │
│  │  (Redis TTL)    │  │     (PostgreSQL tsvector FTS)        │  │
│  └────────┬────────┘  └──────────────────────┬───────────────┘  │
└───────────┼──────────────────────────────────┼──────────────────┘
            │                                  │
┌───────────▼──────────┐         ┌────────────▼──────────────────┐
│      Redis 7         │         │       PostgreSQL 15            │
│  ┌────────────────┐  │         │  ┌──────┐ ┌──────┐ ┌───────┐  │
│  │   API Cache    │  │         │  │ IOCs │ │Actors│ │Bundles│  │
│  │   Rate Limits  │  │         │  └──────┘ └──────┘ └───────┘  │
│  └────────────────┘  │         │  ┌──────────┐ ┌───────────┐   │
└──────────────────────┘         │  │Campaigns │ │ATT&CK Map │   │
                                 │  └──────────┘ └───────────┘   │
                                 └───────────────────────────────┘
```

---

## 📡 API Reference

<details>
<summary><b>🔐 Authentication</b></summary>

```http
POST /api/v1/auth/register
POST /api/v1/auth/login
```

</details>

<details>
<summary><b>🔍 IOC Endpoints</b></summary>

```http
GET    /api/v1/iocs              # List with filters: ?type=ip&tlp=red&q=search
POST   /api/v1/iocs              # Create single IOC
GET    /api/v1/iocs/:id          # Get by ID
PUT    /api/v1/iocs/:id          # Update
DELETE /api/v1/iocs/:id          # Delete
POST   /api/v1/iocs/bulk         # Bulk import array
```

**Filter Parameters:**

| Param | Values | Example |
|-------|--------|---------|
| `type` | ip, domain, url, hash_md5, hash_sha256, email | `?type=ip` |
| `tlp` | white, green, amber, red | `?tlp=red` |
| `tags` | comma-separated | `?tags=apt,russia` |
| `q` | full-text search | `?q=185.220` |
| `page` / `limit` | integer | `?page=2&limit=50` |

</details>

<details>
<summary><b>👤 Threat Actor Endpoints</b></summary>

```http
GET    /api/v1/threat-actors
POST   /api/v1/threat-actors
GET    /api/v1/threat-actors/:id
PUT    /api/v1/threat-actors/:id
DELETE /api/v1/threat-actors/:id
```

</details>

<details>
<summary><b>🎯 Campaign Endpoints</b></summary>

```http
GET    /api/v1/campaigns
POST   /api/v1/campaigns
GET    /api/v1/campaigns/:id
PUT    /api/v1/campaigns/:id
DELETE /api/v1/campaigns/:id
```

</details>

<details>
<summary><b>🗺️ ATT&CK Endpoints</b></summary>

```http
GET    /api/v1/attack/mappings          # List mappings
POST   /api/v1/attack/mappings          # Add mapping
GET    /api/v1/attack/techniques        # Grouped by tactic
DELETE /api/v1/attack/mappings/:id
```

</details>

<details>
<summary><b>📦 STIX / TAXII Endpoints</b></summary>

```http
POST /api/v1/stix/export                # Generate & save STIX 2.1 bundle
POST /api/v1/stix/import                # Parse and import STIX bundle
GET  /taxii/v21/collections             # List TAXII collections
GET  /taxii/v21/collections/:id/objects # Get objects from collection
```

</details>

---

## 🗄️ Data Model

```sql
iocs              threat_actors          campaigns
──────────────    ──────────────────     ──────────────────
id (UUID)    ←─┐  id (UUID)         ←─┐  id (UUID)
type             │  name                │  name
value            │  aliases[]           │  description
tlp_level        │  sophistication      │  first_seen
confidence       │  resource_level      │  last_seen
tags[]           │  primary_motivation  │  objective
source           │  country_code        └─ threat_actor_id ──┐
description      │  description                              │
search_vector    │                                           │
                 │  attack_mappings     stix_bundles         │
                 │  ──────────────────  ──────────────────   │
                 │  id (UUID)           id (UUID)            │
                 └─ entity_id (UUID)    spec_version         │
                    entity_type         bundle_json (JSONB)  │
                    technique_id        created_at           │
                    technique_name                           │
                    tactic                                   │
                    platform[]          ◄────────────────────┘
```

---

## 🎨 UI Screenshots

<div align="center">

| Dashboard | IOC Management |
|-----------|---------------|
| ![Dashboard](https://via.placeholder.com/500x280/0d1117/f78166?text=Dashboard+%E2%80%94+Stats+%2B+Charts) | ![IOCs](https://via.placeholder.com/500x280/0d1117/f78166?text=IOC+Table+%2B+Filters+%2B+TLP+Badges) |

| Threat Actors | ATT&CK Matrix |
|--------------|---------------|
| ![Actors](https://via.placeholder.com/500x280/0d1117/f78166?text=Threat+Actor+Cards+%2B+Attribution) | ![Matrix](https://via.placeholder.com/500x280/0d1117/f78166?text=MITRE+ATT%26CK+Matrix+Overlay) |

</div>

---

## 🧰 Tech Stack

<div align="center">

| Layer | Technology | Purpose |
|-------|-----------|---------|
| **Frontend** | React 18 + TypeScript + Vite | SPA with strict typing |
| **Styling** | Tailwind CSS | Dark theme design system |
| **Data Viz** | Recharts | Confidence & type charts |
| **State** | React Query (@tanstack) | Server state + caching |
| **Backend** | Go 1.21 + Gin | High-performance REST API |
| **Database** | PostgreSQL 15 | Primary data store + FTS |
| **Cache** | Redis 7 | API cache + rate limiting |
| **Auth** | JWT (golang-jwt v5) | Stateless authentication |
| **Standards** | STIX 2.1 / TAXII 2.1 | CTI interoperability |
| **Infra** | Docker Compose | One-command deployment |
| **Migrations** | golang-migrate SQL files | Versioned schema |

</div>

---

## 📁 Project Structure

```
opencti-lite/
│
├── 🐳 docker-compose.yml         # PostgreSQL + Redis
│
├── 🔧 backend/
│   ├── main.go                   # Entry point, router setup
│   ├── go.mod
│   ├── config/
│   │   └── config.go             # Env-based configuration
│   ├── database/
│   │   └── db.go                 # DB + Redis connection
│   ├── migrations/               # 007 versioned SQL files
│   ├── models/                   # IOC, Actor, Campaign, ATT&CK, STIX
│   ├── handlers/                 # HTTP handlers per resource
│   ├── middleware/               # Auth, CORS, Logger, Rate Limit
│   └── services/                 # STIX, Cache, Search
│
└── ⚛️  frontend/
    ├── package.json
    ├── vite.config.ts
    ├── tailwind.config.ts
    └── src/
        ├── types/index.ts        # All TypeScript interfaces
        ├── api/                  # Typed axios API clients
        ├── components/           # Reusable UI components
        │   ├── Layout.tsx
        │   ├── Sidebar.tsx
        │   ├── TLPBadge.tsx
        │   ├── ConfidenceBar.tsx
        │   └── ...
        └── pages/
            ├── Dashboard.tsx
            ├── IOCs.tsx
            ├── ThreatActors.tsx
            ├── Campaigns.tsx
            ├── ATTACKMatrix.tsx
            └── STIXExport.tsx
```

---

## ⚙️ Configuration

```bash
# backend/.env
DB_HOST=localhost
DB_PORT=5432
DB_NAME=opencti
DB_USER=opencti
DB_PASSWORD=opencti
REDIS_ADDR=localhost:6379
JWT_SECRET=change-this-in-production-use-256-bit-key
PORT=8080
```

---

## 🔒 Security

- 🔐 **JWT Authentication** — all `/api/v1/*` routes require valid Bearer token
- 🚦 **Rate Limiting** — 100 requests/minute/IP enforced via Redis
- 🛡️ **CORS** — restricted to configured origins only
- 🔑 **Password Hashing** — bcrypt with cost factor 12
- 📋 **TLP Protocol** — Traffic Light Protocol enforced at data level
- 🏷️ **Role-Based Access** — admin/analyst roles (first user = admin)

---

## 🚢 Production Deployment

<details>
<summary><b>Docker Production Build</b></summary>

```dockerfile
# backend/Dockerfile
FROM golang:1.21-alpine AS builder
WORKDIR /app
COPY . .
RUN go mod tidy && CGO_ENABLED=0 go build -o opencti-lite .

FROM alpine:3.19
RUN apk --no-cache add ca-certificates
COPY --from=builder /app/opencti-lite .
EXPOSE 8080
CMD ["./opencti-lite"]
```

```dockerfile
# frontend/Dockerfile
FROM node:20-alpine AS builder
WORKDIR /app
COPY package*.json .
RUN npm ci
COPY . .
RUN npm run build

FROM nginx:alpine
COPY --from=builder /app/dist /usr/share/nginx/html
EXPOSE 80
```

</details>

<details>
<summary><b>Environment Hardening</b></summary>

```bash
# Generate secure JWT secret
openssl rand -hex 32

# Recommended PostgreSQL production settings
max_connections = 200
shared_buffers = 512MB
effective_cache_size = 2GB
maintenance_work_mem = 128MB

# Redis production settings
maxmemory 512mb
maxmemory-policy allkeys-lru
```

</details>

---

## 🗺️ Roadmap

- [x] IOC Management (6 types, TLP, confidence, bulk import)
- [x] Threat Actor profiling with attribution
- [x] Campaign tracking with timeline view
- [x] MITRE ATT&CK matrix overlay
- [x] STIX 2.1 export / import
- [x] TAXII 2.1 server
- [x] JWT authentication + RBAC
- [x] Redis caching + rate limiting
- [x] Full-text search (PostgreSQL FTS)
- [ ] **v1.1** — Elasticsearch integration for large-scale deployments
- [ ] **v1.2** — Automated IOC feeds (AlienVault OTX, MISP, VirusTotal)
- [ ] **v1.3** — Graph visualization (threat actor → campaign → IOC relationships)
- [ ] **v1.4** — Webhooks & SOAR integration (Splunk, QRadar, Elastic SIEM)
- [ ] **v2.0** — Multi-tenant support + team workspaces

---

## 🤝 Contributing

Contributions are welcome from security researchers, developers, and threat intelligence professionals.

```bash
# 1. Fork the repository
# 2. Create your feature branch
git checkout -b feature/add-misp-integration

# 3. Commit your changes (conventional commits)
git commit -m "feat(feeds): add MISP feed ingestion service"

# 4. Push and open a Pull Request
git push origin feature/add-misp-integration
```

**Areas where help is most needed:**
- 🌐 Additional TAXII feed connectors
- 🧪 Backend unit + integration tests
- 🎨 UI component polish
- 📖 Documentation and tutorials
- 🔌 SIEM/SOAR integration plugins

---

## 📄 License

```
MIT License — Copyright (c) 2024 OpenCTI Lite Contributors

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software to use, copy, modify, merge, publish, distribute, sublicense,
and/or sell copies of the Software, subject to the above copyright notice.
```

---

<div align="center">

<!-- Animated footer -->
<img src="https://capsule-render.vercel.app/api?type=waving&color=0:f78166,100:0d1117&height=120&section=footer&animation=twinkling" width="100%"/>

**Built for security teams who believe threat intelligence should be open.**

<br/>

[![⭐ Star this repo](https://img.shields.io/badge/⭐%20Star%20this%20repo-f78166?style=for-the-badge&labelColor=0d1117)](https://github.com/your-org/opencti-lite)
[![🍴 Fork it](https://img.shields.io/badge/🍴%20Fork%20it-8b949e?style=for-the-badge&labelColor=0d1117)](https://github.com/your-org/opencti-lite/fork)
[![🐛 Report Bug](https://img.shields.io/badge/🐛%20Report%20Bug-f85149?style=for-the-badge&labelColor=0d1117)](https://github.com/your-org/opencti-lite/issues)

<br/>

*If this project saves your team money, consider starring it ⭐*

</div>
