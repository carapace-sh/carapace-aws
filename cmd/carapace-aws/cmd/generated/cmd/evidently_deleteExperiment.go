package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_deleteExperimentCmd = &cobra.Command{
	Use:   "delete-experiment",
	Short: "Deletes an Evidently experiment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_deleteExperimentCmd).Standalone()

	evidently_deleteExperimentCmd.Flags().String("experiment", "", "The name of the experiment to delete.")
	evidently_deleteExperimentCmd.Flags().String("project", "", "The name or ARN of the project that contains the experiment to delete.")
	evidently_deleteExperimentCmd.MarkFlagRequired("experiment")
	evidently_deleteExperimentCmd.MarkFlagRequired("project")
	evidentlyCmd.AddCommand(evidently_deleteExperimentCmd)
}
