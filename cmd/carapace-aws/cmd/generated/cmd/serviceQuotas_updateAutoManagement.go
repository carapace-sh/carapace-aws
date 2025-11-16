package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var serviceQuotas_updateAutoManagementCmd = &cobra.Command{
	Use:   "update-auto-management",
	Short: "Updates your [Service Quotas Automatic Management](https://docs.aws.amazon.com/servicequotas/latest/userguide/automatic-management.html) configuration, including notification preferences and excluded quotas.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(serviceQuotas_updateAutoManagementCmd).Standalone()

	serviceQuotas_updateAutoManagementCmd.Flags().String("exclusion-list", "", "List of Amazon Web Services services you want to exclude from Automatic Management.")
	serviceQuotas_updateAutoManagementCmd.Flags().String("notification-arn", "", "The [User Notifications](https://docs.aws.amazon.com/notifications/latest/userguide/resource-level-permissions.html#rlp-table) Amazon Resource Name (ARN) for Automatic Management notifications you want to update.")
	serviceQuotas_updateAutoManagementCmd.Flags().String("opt-in-type", "", "Information on the opt-in type for your Automatic Management configuration.")
	serviceQuotasCmd.AddCommand(serviceQuotas_updateAutoManagementCmd)
}
