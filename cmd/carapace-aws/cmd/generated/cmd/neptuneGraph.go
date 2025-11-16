package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var neptuneGraphCmd = &cobra.Command{
	Use:   "neptune-graph",
	Short: "Neptune Analytics is a new analytics database engine for Amazon Neptune that helps customers get to insights faster by quickly processing large amounts of graph data, invoking popular graph analytic algorithms in low-latency queries, and getting analytics results in seconds.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(neptuneGraphCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(neptuneGraphCmd).Standalone()

	})
	rootCmd.AddCommand(neptuneGraphCmd)
}
