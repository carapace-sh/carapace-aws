package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fis_startExperimentCmd = &cobra.Command{
	Use:   "start-experiment",
	Short: "Starts running an experiment from the specified experiment template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fis_startExperimentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fis_startExperimentCmd).Standalone()

		fis_startExperimentCmd.Flags().String("client-token", "", "Unique, case-sensitive identifier that you provide to ensure the idempotency of the request.")
		fis_startExperimentCmd.Flags().String("experiment-options", "", "The experiment options for running the experiment.")
		fis_startExperimentCmd.Flags().String("experiment-template-id", "", "The ID of the experiment template.")
		fis_startExperimentCmd.Flags().String("tags", "", "The tags to apply to the experiment.")
		fis_startExperimentCmd.MarkFlagRequired("client-token")
		fis_startExperimentCmd.MarkFlagRequired("experiment-template-id")
	})
	fisCmd.AddCommand(fis_startExperimentCmd)
}
