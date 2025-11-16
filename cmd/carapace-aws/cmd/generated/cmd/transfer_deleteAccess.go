package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var transfer_deleteAccessCmd = &cobra.Command{
	Use:   "delete-access",
	Short: "Allows you to delete the access specified in the `ServerID` and `ExternalID` parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(transfer_deleteAccessCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(transfer_deleteAccessCmd).Standalone()

		transfer_deleteAccessCmd.Flags().String("external-id", "", "A unique identifier that is required to identify specific groups within your directory.")
		transfer_deleteAccessCmd.Flags().String("server-id", "", "A system-assigned unique identifier for a server that has this user assigned.")
		transfer_deleteAccessCmd.MarkFlagRequired("external-id")
		transfer_deleteAccessCmd.MarkFlagRequired("server-id")
	})
	transferCmd.AddCommand(transfer_deleteAccessCmd)
}
