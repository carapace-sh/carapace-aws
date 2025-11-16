package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discovery_startContinuousExportCmd = &cobra.Command{
	Use:   "start-continuous-export",
	Short: "Start the continuous flow of agent's discovered data into Amazon Athena.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discovery_startContinuousExportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(discovery_startContinuousExportCmd).Standalone()

	})
	discoveryCmd.AddCommand(discovery_startContinuousExportCmd)
}
