---
"fingerprint-pro-server-api-go-sdk": patch
---

Percent-encode path parameter values so they remain within a single URL path segment.

`GetEvent`, `UpdateEvent`, `GetVisits`, and `DeleteVisitorData` now return the new `InvalidArgumentError` without sending a request when their identifier argument is `.` or `..`. Those values are not valid identifiers. The error exposes the rejected `Parameter()` and `Value()`.
