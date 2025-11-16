package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_startExperimentCmd = &cobra.Command{
	Use:   "start-experiment",
	Short: "Starts an existing experiment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_startExperimentCmd).Standalone()

	evidently_startExperimentCmd.Flags().String("analysis-complete-time", "", "The date and time to end the experiment.")
	evidently_startExperimentCmd.Flags().String("experiment", "", "The name of the experiment to start.")
	evidently_startExperimentCmd.Flags().String("project", "", "The name or ARN of the project that contains the experiment to start.")
	evidently_startExperimentCmd.MarkFlagRequired("analysis-complete-time")
	evidently_startExperimentCmd.MarkFlagRequired("experiment")
	evidently_startExperimentCmd.MarkFlagRequired("project")
	evidentlyCmd.AddCommand(evidently_startExperimentCmd)
}
