package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediaconnect_listFlowsCmd = &cobra.Command{
	Use:   "list-flows",
	Short: "Displays a list of flows that are associated with this account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediaconnect_listFlowsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediaconnect_listFlowsCmd).Standalone()

		mediaconnect_listFlowsCmd.Flags().String("max-results", "", "The maximum number of results to return per API request.")
		mediaconnect_listFlowsCmd.Flags().String("next-token", "", "The token that identifies the batch of results that you want to see.")
	})
	mediaconnectCmd.AddCommand(mediaconnect_listFlowsCmd)
}
