---
layout: "loopia"
page_title: "Provider: Loopia"
description: |-
  The Loopia provider is used to interact with the resources supported by Loopia. The provider needs to be configured with the proper credentials before it can be used.
---

# Loopia Provider

## Example Usage

```hcl
# Configure the Loopia provider
provider "loopia" {
  username = "${var.loopia_username}"
  password = "${var.loopia_password}"
}

# Create a zone record
resource "loopia_record" "default" {
  # ...
}
```
