package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fis_stopExperimentCmd = &cobra.Command{
	Use:   "stop-experiment",
	Short: "Stops the specified experiment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fis_stopExperimentCmd).Standalone()

	fis_stopExperimentCmd.Flags().String("id", "", "The ID of the experiment.")
	fis_stopExperimentCmd.MarkFlagRequired("id")
	fisCmd.AddCommand(fis_stopExperimentCmd)
}
