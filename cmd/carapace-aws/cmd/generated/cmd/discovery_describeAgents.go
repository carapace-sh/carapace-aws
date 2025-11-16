package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discovery_describeAgentsCmd = &cobra.Command{
	Use:   "describe-agents",
	Short: "Lists agents or collectors as specified by ID or other filters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discovery_describeAgentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(discovery_describeAgentsCmd).Standalone()

		discovery_describeAgentsCmd.Flags().String("agent-ids", "", "The agent or the collector IDs for which you want information.")
		discovery_describeAgentsCmd.Flags().String("filters", "", "You can filter the request using various logical operators and a *key*-*value* format.")
		discovery_describeAgentsCmd.Flags().String("max-results", "", "The total number of agents/collectors to return in a single page of output.")
		discovery_describeAgentsCmd.Flags().String("next-token", "", "Token to retrieve the next set of results.")
	})
	discoveryCmd.AddCommand(discovery_describeAgentsCmd)
}
