package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var discovery_stopContinuousExportCmd = &cobra.Command{
	Use:   "stop-continuous-export",
	Short: "Stop the continuous flow of agent's discovered data into Amazon Athena.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(discovery_stopContinuousExportCmd).Standalone()

	discovery_stopContinuousExportCmd.Flags().String("export-id", "", "The unique ID assigned to this export.")
	discovery_stopContinuousExportCmd.MarkFlagRequired("export-id")
	discoveryCmd.AddCommand(discovery_stopContinuousExportCmd)
}
