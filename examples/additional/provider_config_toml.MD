[example]
account_name = 'account_name'
organization_name = 'organization_name'
user = 'user'
password = 'password'
warehouse = 'SNOWFLAKE'
role = 'ACCOUNTADMIN'
client_ip = '1.2.3.4'
protocol = 'https'
port = 443
okta_url = 'https://example.com'
client_timeout = 10
jwt_client_timeout = 20
login_timeout = 30
request_timeout = 40
jwt_expire_timeout = 50
external_browser_timeout = 60
max_retry_count = 1
authenticator = 'snowflake'
insecure_mode = true
ocsp_fail_open = true
keep_session_alive = true
disable_telemetry = true
validate_default_parameters = true
client_request_mfa_token = true
client_store_temporary_credential = true
driver_tracing = 'info'
tmp_dir_path = '/tmp/terraform-provider/'
disable_query_context_cache = true
include_retry_reason = true
disable_console_login = true

[example.params]
param_key = 'param_value'
