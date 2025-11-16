package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteActionConnectorCmd = &cobra.Command{
	Use:   "delete-action-connector",
	Short: "Hard deletes an action connector, making it unrecoverable.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteActionConnectorCmd).Standalone()

	quicksight_deleteActionConnectorCmd.Flags().String("action-connector-id", "", "The unique identifier of the action connector to delete.")
	quicksight_deleteActionConnectorCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID that contains the action connector to delete.")
	quicksight_deleteActionConnectorCmd.MarkFlagRequired("action-connector-id")
	quicksight_deleteActionConnectorCmd.MarkFlagRequired("aws-account-id")
	quicksightCmd.AddCommand(quicksight_deleteActionConnectorCmd)
}
