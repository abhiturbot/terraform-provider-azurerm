package resource_test

import (
	"fmt"
	"github.com/hashicorp/terraform-provider-azurerm/internal/acceptance"
)

type ManagementGroupTemplateDeploymentDataSource struct {
}

func (ManagementGroupTemplateDeploymentDataSource) withOutputsConfig(data acceptance.TestData) string {
	return fmt.Sprintf(`
provider "azurerm" {
  features {}
}

resource "azurerm_management_group" "test" {
  name = "TestAcc-Deployment-%[1]d"
}

resource "azurerm_management_group_template_deployment" "test" {
  name                = "acctestMGdeploy-%[1]d"
  management_group_id = azurerm_management_group.test.id
  location            = %[2]q

  template_content = <<TEMPLATE
{
 "$schema": "https://schema.management.azure.com/schemas/2015-01-01/deploymentTemplate.json#",
 "contentVersion": "1.0.0.0",
 "parameters": {},
 "variables": {},
 "resources": [],
 "outputs": {
    "testOutput": {
      "type": "String",
      "value": "some-value"
    }
  }
}
TEMPLATE
}
`, data.RandomInteger, data.Locations.Primary)
}

func (r ManagementGroupTemplateDeploymentDataSource) withDataSource(data acceptance.TestData) string {
	return fmt.Sprintf(`
%s

data "azurerm_management_group_template_deployment" "test" {
  name                  = azurerm_management_group_template_deployment.test.name
  management_group_name = azurerm_management_group.test.name
}
`, r.withOutputsConfig(data))
}
