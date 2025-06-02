---
"fingerprint-pro-server-api-go-sdk": minor
---

Added explicit enum constant prefixes for better clarity and to avoid name collisions.
The following existing enums now use prefixed constant names (all new enums will follow this convention):
  - `VpnConfidence`
  - `ErrorCode`
  - `BotdBotResult`
Deprecated the old unprefixed constant aliases; they will be removed in the next major release.
