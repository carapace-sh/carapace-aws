package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeActionConnectorPermissionsCmd = &cobra.Command{
	Use:   "describe-action-connector-permissions",
	Short: "Retrieves the permissions configuration for an action connector, showing which users, groups, and namespaces have access and what operations they can perform.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeActionConnectorPermissionsCmd).Standalone()

	quicksight_describeActionConnectorPermissionsCmd.Flags().String("action-connector-id", "", "The unique identifier of the action connector whose permissions you want to describe.")
	quicksight_describeActionConnectorPermissionsCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID that contains the action connector.")
	quicksight_describeActionConnectorPermissionsCmd.MarkFlagRequired("action-connector-id")
	quicksight_describeActionConnectorPermissionsCmd.MarkFlagRequired("aws-account-id")
	quicksightCmd.AddCommand(quicksight_describeActionConnectorPermissionsCmd)
}
