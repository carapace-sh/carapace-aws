package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeActionConnectorCmd = &cobra.Command{
	Use:   "describe-action-connector",
	Short: "Retrieves detailed information about an action connector, including its configuration, authentication settings, enabled actions, and current status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeActionConnectorCmd).Standalone()

	quicksight_describeActionConnectorCmd.Flags().String("action-connector-id", "", "The unique identifier of the action connector to describe.")
	quicksight_describeActionConnectorCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID that contains the action connector.")
	quicksight_describeActionConnectorCmd.MarkFlagRequired("action-connector-id")
	quicksight_describeActionConnectorCmd.MarkFlagRequired("aws-account-id")
	quicksightCmd.AddCommand(quicksight_describeActionConnectorCmd)
}
