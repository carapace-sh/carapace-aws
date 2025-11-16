package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_confirmConnectionCmd = &cobra.Command{
	Use:   "confirm-connection",
	Short: "Confirms the creation of the specified hosted connection on an interconnect.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_confirmConnectionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(directconnect_confirmConnectionCmd).Standalone()

		directconnect_confirmConnectionCmd.Flags().String("connection-id", "", "The ID of the hosted connection.")
		directconnect_confirmConnectionCmd.MarkFlagRequired("connection-id")
	})
	directconnectCmd.AddCommand(directconnect_confirmConnectionCmd)
}
