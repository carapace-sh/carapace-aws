package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_deleteConnectorCmd = &cobra.Command{
	Use:   "delete-connector",
	Short: "Deletes the connector that's specified in the provided `ConnectorId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_deleteConnectorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transfer_deleteConnectorCmd).Standalone()

		transfer_deleteConnectorCmd.Flags().String("connector-id", "", "The unique identifier for the connector.")
		transfer_deleteConnectorCmd.MarkFlagRequired("connector-id")
	})
	transferCmd.AddCommand(transfer_deleteConnectorCmd)
}
