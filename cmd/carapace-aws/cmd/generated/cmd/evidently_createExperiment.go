package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_createExperimentCmd = &cobra.Command{
	Use:   "create-experiment",
	Short: "Creates an Evidently *experiment*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_createExperimentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(evidently_createExperimentCmd).Standalone()

		evidently_createExperimentCmd.Flags().String("description", "", "An optional description of the experiment.")
		evidently_createExperimentCmd.Flags().String("metric-goals", "", "An array of structures that defines the metrics used for the experiment, and whether a higher or lower value for each metric is the goal.")
		evidently_createExperimentCmd.Flags().String("name", "", "A name for the new experiment.")
		evidently_createExperimentCmd.Flags().String("online-ab-config", "", "A structure that contains the configuration of which variation to use as the \"control\" version.")
		evidently_createExperimentCmd.Flags().String("project", "", "The name or ARN of the project that you want to create the new experiment in.")
		evidently_createExperimentCmd.Flags().String("randomization-salt", "", "When Evidently assigns a particular user session to an experiment, it must use a randomization ID to determine which variation the user session is served.")
		evidently_createExperimentCmd.Flags().String("sampling-rate", "", "The portion of the available audience that you want to allocate to this experiment, in thousandths of a percent.")
		evidently_createExperimentCmd.Flags().String("segment", "", "Specifies an audience *segment* to use in the experiment.")
		evidently_createExperimentCmd.Flags().String("tags", "", "Assigns one or more tags (key-value pairs) to the experiment.")
		evidently_createExperimentCmd.Flags().String("treatments", "", "An array of structures that describe the configuration of each feature variation used in the experiment.")
		evidently_createExperimentCmd.MarkFlagRequired("metric-goals")
		evidently_createExperimentCmd.MarkFlagRequired("name")
		evidently_createExperimentCmd.MarkFlagRequired("project")
		evidently_createExperimentCmd.MarkFlagRequired("treatments")
	})
	evidentlyCmd.AddCommand(evidently_createExperimentCmd)
}
