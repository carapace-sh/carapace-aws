package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discovery_getDiscoverySummaryCmd = &cobra.Command{
	Use:   "get-discovery-summary",
	Short: "Retrieves a short summary of discovered assets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discovery_getDiscoverySummaryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(discovery_getDiscoverySummaryCmd).Standalone()

	})
	discoveryCmd.AddCommand(discovery_getDiscoverySummaryCmd)
}
