package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_stopExperimentCmd = &cobra.Command{
	Use:   "stop-experiment",
	Short: "Stops an experiment that is currently running.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_stopExperimentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(evidently_stopExperimentCmd).Standalone()

		evidently_stopExperimentCmd.Flags().String("desired-state", "", "Specify whether the experiment is to be considered `COMPLETED` or `CANCELLED` after it stops.")
		evidently_stopExperimentCmd.Flags().String("experiment", "", "The name of the experiment to stop.")
		evidently_stopExperimentCmd.Flags().String("project", "", "The name or ARN of the project that contains the experiment to stop.")
		evidently_stopExperimentCmd.Flags().String("reason", "", "A string that describes why you are stopping the experiment.")
		evidently_stopExperimentCmd.MarkFlagRequired("experiment")
		evidently_stopExperimentCmd.MarkFlagRequired("project")
	})
	evidentlyCmd.AddCommand(evidently_stopExperimentCmd)
}
