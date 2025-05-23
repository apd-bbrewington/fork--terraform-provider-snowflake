---
page_title: "snowflake_secondary_connection Resource - terraform-provider-snowflake"
subcategory: "Stable"
description: |-
  Resource used to manage secondary (replicated) connections. To manage primary connection check resource snowflake_primary_connection ./primary_connection. For more information, check connection documentation https://docs.snowflake.com/en/sql-reference/sql/create-connection.html.
---

# snowflake_secondary_connection (Resource)

Resource used to manage secondary (replicated) connections. To manage primary connection check resource [snowflake_primary_connection](./primary_connection). For more information, check [connection documentation](https://docs.snowflake.com/en/sql-reference/sql/create-connection.html).

## Example Usage

```terraform
## Minimal
resource "snowflake_secondary_connection" "basic" {
  name          = "connection_name"
  as_replica_of = "\"<organization_name>\".\"<account_name>\".\"<connection_name>\""
}

## Complete (with every optional set)
resource "snowflake_secondary_connection" "complete" {
  name          = "connection_name"
  as_replica_of = "\"<organization_name>\".\"<account_name>\".\"<connection_name>\""
  comment       = "my complete secondary connection"
}
```

-> **Note** Instead of using fully_qualified_name, you can reference objects managed outside Terraform by constructing a correct ID, consult [identifiers guide](../guides/identifiers_rework_design_decisions#new-computed-fully-qualified-name-field-in-resources).

-> **Note** To promote `snowflake_secondary_connection` to [`snowflake_primary_connection`](./primary_connection), resources need to be migrated manually. For guidance on removing and importing resources into the state check [resource migration](../guides/resource_migration). Remove the resource from the state with [terraform state rm](https://developer.hashicorp.com/terraform/cli/commands/state/rm), then promote it manually using:
    ```
    ALTER CONNECTION <name> PRIMARY;
    ```
and then import it as the `snowflake_primary_connection`.
<!-- TODO(SNOW-1634854): include an example showing both methods-->

-> **Note** If a field has a default value, it is shown next to the type in the schema.

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `as_replica_of` (String) Specifies the identifier for a primary connection from which to create a replica (i.e. a secondary connection). For more information about this resource, see [docs](./primary_connection).
- `name` (String) String that specifies the identifier (i.e. name) for the connection. Must start with an alphabetic character and may only contain letters, decimal digits (0-9), and underscores (_). For a secondary connection, the name must match the name of its primary connection. Due to technical limitations (read more [here](../guides/identifiers_rework_design_decisions#known-limitations-and-identifier-recommendations)), avoid using the following characters: `|`, `.`, `"`.

### Optional

- `comment` (String) Specifies a comment for the secondary connection.
- `timeouts` (Block, Optional) (see [below for nested schema](#nestedblock--timeouts))

### Read-Only

- `fully_qualified_name` (String) Fully qualified name of the resource. For more information, see [object name resolution](https://docs.snowflake.com/en/sql-reference/name-resolution).
- `id` (String) The ID of this resource.
- `is_primary` (Boolean) Indicates if the connection primary status has been changed. If change is detected, resource will be recreated.
- `show_output` (List of Object) Outputs the result of `SHOW CONNECTIONS` for the given connection. (see [below for nested schema](#nestedatt--show_output))

<a id="nestedblock--timeouts"></a>
### Nested Schema for `timeouts`

Optional:

- `create` (String)
- `delete` (String)
- `read` (String)
- `update` (String)


<a id="nestedatt--show_output"></a>
### Nested Schema for `show_output`

Read-Only:

- `account_locator` (String)
- `account_name` (String)
- `comment` (String)
- `connection_url` (String)
- `created_on` (String)
- `failover_allowed_to_accounts` (List of String)
- `is_primary` (Boolean)
- `name` (String)
- `organization_name` (String)
- `primary` (String)
- `region_group` (String)
- `snowflake_region` (String)

## Import

Import is supported using the following syntax:

```shell
terraform import snowflake_secondary_connection.example '"<secondary_connection_name>"'
```
