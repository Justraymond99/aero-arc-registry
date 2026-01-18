# Aero Arc Relay Registry

## Overview
The Aero Arc Relay Registry is a control-plane gRPC service for coordinating live Aero Arc Relay instances and their agent ownership. It provides a single, backend-agnostic source of current routing metadata so API servers and operator dashboards can make decisions without coupling to any specific storage system.

This solves the problem of keeping relay liveness and agent ownership visible and consistent enough for routing, while remaining replaceable and tolerant of backend failures.

## Architecture
- **Control-plane service**: stores and serves metadata only.
- **Data-plane**: relay and agent traffic flows elsewhere; the registry never forwards traffic.
- **gRPC-only**: all external interaction happens over gRPC.
- **Pluggable storage backends**: Redis, Consul, etcd, and in-memory implementations are supported through a Go interface.

In the broader Aero Arc system, the registry sits between relays/agents and control-plane consumers. Relays and agents register and renew TTL-based ownership; control-plane consumers query the current state to drive routing and operational views.

## Design Goals
- Keep the service simple, readable, and replaceable.
- Maintain clear separation of control-plane metadata from data-plane traffic.
- Stay backend-agnostic; do not leak backend concepts into the API.
- Rely on TTL-based liveness and ownership with eventual consistency.
- Degrade gracefully when backends fail.

## Non-Goals
- Data-plane traffic forwarding or routing decisions.
- Strong consistency, global ordering, or consensus guarantees.
- Backend-specific APIs or operational dependencies exposed to clients.
- Historical analytics or long-term state retention.

## High-Level API
- Register and renew relay liveness (TTL-based).
- Register and renew agent-to-relay ownership (TTL-based).
- Query current relay and ownership state for routing and operator views.

## Status / Roadmap
- Early, focused control-plane service with a stable gRPC surface.
- Backend implementations and operational tooling will evolve independently.
- Backward-compatible API evolution is prioritized over feature expansion.
