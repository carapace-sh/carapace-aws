package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_deleteInterconnectCmd = &cobra.Command{
	Use:   "delete-interconnect",
	Short: "Deletes the specified interconnect.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_deleteInterconnectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(directconnect_deleteInterconnectCmd).Standalone()

		directconnect_deleteInterconnectCmd.Flags().String("interconnect-id", "", "The ID of the interconnect.")
		directconnect_deleteInterconnectCmd.MarkFlagRequired("interconnect-id")
	})
	directconnectCmd.AddCommand(directconnect_deleteInterconnectCmd)
}
