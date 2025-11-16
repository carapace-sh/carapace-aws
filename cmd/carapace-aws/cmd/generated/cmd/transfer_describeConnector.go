package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_describeConnectorCmd = &cobra.Command{
	Use:   "describe-connector",
	Short: "Describes the connector that's identified by the `ConnectorId.`",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_describeConnectorCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transfer_describeConnectorCmd).Standalone()

		transfer_describeConnectorCmd.Flags().String("connector-id", "", "The unique identifier for the connector.")
		transfer_describeConnectorCmd.MarkFlagRequired("connector-id")
	})
	transferCmd.AddCommand(transfer_describeConnectorCmd)
}
