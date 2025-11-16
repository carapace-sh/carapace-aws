package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_updateGlobalTableSettingsCmd = &cobra.Command{
	Use:   "update-global-table-settings",
	Short: "Updates settings for a global table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_updateGlobalTableSettingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dynamodb_updateGlobalTableSettingsCmd).Standalone()

		dynamodb_updateGlobalTableSettingsCmd.Flags().String("global-table-billing-mode", "", "The billing mode of the global table.")
		dynamodb_updateGlobalTableSettingsCmd.Flags().String("global-table-global-secondary-index-settings-update", "", "Represents the settings of a global secondary index for a global table that will be modified.")
		dynamodb_updateGlobalTableSettingsCmd.Flags().String("global-table-name", "", "The name of the global table")
		dynamodb_updateGlobalTableSettingsCmd.Flags().String("global-table-provisioned-write-capacity-auto-scaling-settings-update", "", "Auto scaling settings for managing provisioned write capacity for the global table.")
		dynamodb_updateGlobalTableSettingsCmd.Flags().String("global-table-provisioned-write-capacity-units", "", "The maximum number of writes consumed per second before DynamoDB returns a `ThrottlingException.`")
		dynamodb_updateGlobalTableSettingsCmd.Flags().String("replica-settings-update", "", "Represents the settings for a global table in a Region that will be modified.")
		dynamodb_updateGlobalTableSettingsCmd.MarkFlagRequired("global-table-name")
	})
	dynamodbCmd.AddCommand(dynamodb_updateGlobalTableSettingsCmd)
}
