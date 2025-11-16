package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_updateActionConnectorCmd = &cobra.Command{
	Use:   "update-action-connector",
	Short: "Updates an existing action connector with new configuration details, authentication settings, or enabled actions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_updateActionConnectorCmd).Standalone()

	quicksight_updateActionConnectorCmd.Flags().String("action-connector-id", "", "The unique identifier of the action connector to update.")
	quicksight_updateActionConnectorCmd.Flags().String("authentication-config", "", "The updated authentication configuration for connecting to the external service.")
	quicksight_updateActionConnectorCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID that contains the action connector to update.")
	quicksight_updateActionConnectorCmd.Flags().String("description", "", "The updated description of the action connector.")
	quicksight_updateActionConnectorCmd.Flags().String("name", "", "The new name for the action connector.")
	quicksight_updateActionConnectorCmd.Flags().String("vpc-connection-arn", "", "The updated ARN of the VPC connection to use for secure connectivity.")
	quicksight_updateActionConnectorCmd.MarkFlagRequired("action-connector-id")
	quicksight_updateActionConnectorCmd.MarkFlagRequired("authentication-config")
	quicksight_updateActionConnectorCmd.MarkFlagRequired("aws-account-id")
	quicksight_updateActionConnectorCmd.MarkFlagRequired("name")
	quicksightCmd.AddCommand(quicksight_updateActionConnectorCmd)
}
