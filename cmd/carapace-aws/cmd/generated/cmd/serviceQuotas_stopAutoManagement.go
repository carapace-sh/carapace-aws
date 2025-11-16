package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serviceQuotas_stopAutoManagementCmd = &cobra.Command{
	Use:   "stop-auto-management",
	Short: "Stops [Service Quotas Automatic Management](https://docs.aws.amazon.com/servicequotas/latest/userguide/automatic-management.html) for an Amazon Web Services account and removes all associated configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceQuotas_stopAutoManagementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(serviceQuotas_stopAutoManagementCmd).Standalone()

	})
	serviceQuotasCmd.AddCommand(serviceQuotas_stopAutoManagementCmd)
}
