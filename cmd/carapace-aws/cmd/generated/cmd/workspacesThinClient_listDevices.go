package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var workspacesThinClient_listDevicesCmd = &cobra.Command{
	Use:   "list-devices",
	Short: "Returns a list of thin client devices.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(workspacesThinClient_listDevicesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(workspacesThinClient_listDevicesCmd).Standalone()

		workspacesThinClient_listDevicesCmd.Flags().String("max-results", "", "The maximum number of results that are returned per call.")
		workspacesThinClient_listDevicesCmd.Flags().String("next-token", "", "If `nextToken` is returned, there are more results available.")
	})
	workspacesThinClientCmd.AddCommand(workspacesThinClient_listDevicesCmd)
}
