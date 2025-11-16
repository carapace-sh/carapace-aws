package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serviceQuotas_startAutoManagementCmd = &cobra.Command{
	Use:   "start-auto-management",
	Short: "Starts [Service Quotas Automatic Management](https://docs.aws.amazon.com/servicequotas/latest/userguide/automatic-management.html) for an Amazon Web Services account, including notification preferences and excluded quotas configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceQuotas_startAutoManagementCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(serviceQuotas_startAutoManagementCmd).Standalone()

		serviceQuotas_startAutoManagementCmd.Flags().String("exclusion-list", "", "List of Amazon Web Services services excluded from Automatic Management.")
		serviceQuotas_startAutoManagementCmd.Flags().String("notification-arn", "", "The [User Notifications](https://docs.aws.amazon.com/notifications/latest/userguide/resource-level-permissions.html#rlp-table) Amazon Resource Name (ARN) for Automatic Management notifications.")
		serviceQuotas_startAutoManagementCmd.Flags().String("opt-in-level", "", "Sets the opt-in level for Automatic Management.")
		serviceQuotas_startAutoManagementCmd.Flags().String("opt-in-type", "", "Sets the opt-in type for Automatic Management.")
		serviceQuotas_startAutoManagementCmd.MarkFlagRequired("opt-in-level")
		serviceQuotas_startAutoManagementCmd.MarkFlagRequired("opt-in-type")
	})
	serviceQuotasCmd.AddCommand(serviceQuotas_startAutoManagementCmd)
}
