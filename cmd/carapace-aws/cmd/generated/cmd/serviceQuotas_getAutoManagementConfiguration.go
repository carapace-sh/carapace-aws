package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serviceQuotas_getAutoManagementConfigurationCmd = &cobra.Command{
	Use:   "get-auto-management-configuration",
	Short: "Retrieves information about your [Service Quotas Automatic Management](https://docs.aws.amazon.com/servicequotas/latest/userguide/automatic-management.html) configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceQuotas_getAutoManagementConfigurationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(serviceQuotas_getAutoManagementConfigurationCmd).Standalone()

	})
	serviceQuotasCmd.AddCommand(serviceQuotas_getAutoManagementConfigurationCmd)
}
