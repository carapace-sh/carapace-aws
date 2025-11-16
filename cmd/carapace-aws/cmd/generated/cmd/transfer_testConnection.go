package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_testConnectionCmd = &cobra.Command{
	Use:   "test-connection",
	Short: "Tests whether your SFTP connector is set up successfully.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_testConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transfer_testConnectionCmd).Standalone()

		transfer_testConnectionCmd.Flags().String("connector-id", "", "The unique identifier for the connector.")
		transfer_testConnectionCmd.MarkFlagRequired("connector-id")
	})
	transferCmd.AddCommand(transfer_testConnectionCmd)
}
