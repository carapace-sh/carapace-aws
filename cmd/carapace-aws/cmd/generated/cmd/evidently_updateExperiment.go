package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_updateExperimentCmd = &cobra.Command{
	Use:   "update-experiment",
	Short: "Updates an Evidently experiment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_updateExperimentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(evidently_updateExperimentCmd).Standalone()

		evidently_updateExperimentCmd.Flags().String("description", "", "An optional description of the experiment.")
		evidently_updateExperimentCmd.Flags().String("experiment", "", "The name of the experiment to update.")
		evidently_updateExperimentCmd.Flags().String("metric-goals", "", "An array of structures that defines the metrics used for the experiment, and whether a higher or lower value for each metric is the goal.")
		evidently_updateExperimentCmd.Flags().String("online-ab-config", "", "A structure that contains the configuration of which variation o use as the \"control\" version.")
		evidently_updateExperimentCmd.Flags().String("project", "", "The name or ARN of the project that contains the experiment that you want to update.")
		evidently_updateExperimentCmd.Flags().String("randomization-salt", "", "When Evidently assigns a particular user session to an experiment, it must use a randomization ID to determine which variation the user session is served.")
		evidently_updateExperimentCmd.Flags().String("remove-segment", "", "Removes a segment from being used in an experiment.")
		evidently_updateExperimentCmd.Flags().String("sampling-rate", "", "The portion of the available audience that you want to allocate to this experiment, in thousandths of a percent.")
		evidently_updateExperimentCmd.Flags().String("segment", "", "Adds an audience *segment* to an experiment.")
		evidently_updateExperimentCmd.Flags().String("treatments", "", "An array of structures that define the variations being tested in the experiment.")
		evidently_updateExperimentCmd.MarkFlagRequired("experiment")
		evidently_updateExperimentCmd.MarkFlagRequired("project")
	})
	evidentlyCmd.AddCommand(evidently_updateExperimentCmd)
}
