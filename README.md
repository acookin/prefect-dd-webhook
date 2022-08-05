# Prefect webhook client

A server to recieve webhook request from [Prefect](https://www.prefect.io/) and
transform then relay those requests as datadog events.

This is meant to be used with the
[cloud hooks](https://docs-v1.prefect.io/orchestration/concepts/cloud_hooks.html#ui) in
prefect 1.0.

At time of writing, an equivalent webhook feature has not been released for prefect 2.0.


## Usage

1. Deploy the image to the container management system of your choosing
2. Set `DD_API_KEY` to an api key for your DD account
(https://docs.datadoghq.com/account_management/api-app-keys/) as an environment
variable.
3. Setup ingress/dns to the webhook client service
4. Add the webhook hostname as a cloud webhook in prefect
   https://docs-v1.prefect.io/orchestration/concepts/cloud_hooks.html#ui
    * you can filter the types of flow events you want to send.
5. Watch as events start trickling into datadog
