# Prefect webhook client

A server to recieve webhook request from [Prefect](https://www.prefect.io/) and
transform then relay those requests as datadog events.

This is meant to be used with the
[cloud hooks](https://docs-v1.prefect.io/orchestration/concepts/cloud_hooks.html#ui) in
prefect 1.0.

At time of writing, an equivalent webhook feature has not been released for prefect 2.0.
