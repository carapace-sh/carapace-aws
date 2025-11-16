package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var directconnect_describeConnectionsCmd = &cobra.Command{
	Use:   "describe-connections",
	Short: "Displays the specified connection or all connections in this Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(directconnect_describeConnectionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(directconnect_describeConnectionsCmd).Standalone()

		directconnect_describeConnectionsCmd.Flags().String("connection-id", "", "The ID of the connection.")
		directconnect_describeConnectionsCmd.Flags().String("max-results", "", "The maximum number of results to return with a single call.")
		directconnect_describeConnectionsCmd.Flags().String("next-token", "", "The token for the next page of results.")
	})
	directconnectCmd.AddCommand(directconnect_describeConnectionsCmd)
}
