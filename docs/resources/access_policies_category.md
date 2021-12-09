---
# generated by https://github.com/hashicorp/terraform-plugin-docs
page_title: "fmc_access_policies_category Resource - terraform-provider-fmc"
subcategory: ""
description: |-
  Resource for access policy category in FMC
  Example
  An example is shown below:
  hcl
  resource "fmc_access_policies_category" "category" {
      name                  = "test-time-range"
      access_policy_id     = "BB62F664-7168-4C8E-B4CE-F70D522889D2"
  }
---

# fmc_access_policies_category (Resource)

Resource for access policy category in FMC

## Example
An example is shown below: 
```hcl
resource "fmc_access_policies_category" "category" {
    name        		  = "test-time-range"
    access_policy_id     = "BB62F664-7168-4C8E-B4CE-F70D522889D2"
}
```



<!-- schema generated by tfplugindocs -->
## Schema

### Required

- **access_policy_id** (String) Id of access policy this category belongs to
- **name** (String) The name of this category

### Optional

- **id** (String) The ID of this resource.

