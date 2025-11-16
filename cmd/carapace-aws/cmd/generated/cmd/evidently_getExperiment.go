package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var evidently_getExperimentCmd = &cobra.Command{
	Use:   "get-experiment",
	Short: "Returns the details about one experiment.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(evidently_getExperimentCmd).Standalone()

	evidently_getExperimentCmd.Flags().String("experiment", "", "The name of the experiment that you want to see the details of.")
	evidently_getExperimentCmd.Flags().String("project", "", "The name or ARN of the project that contains the experiment.")
	evidently_getExperimentCmd.MarkFlagRequired("experiment")
	evidently_getExperimentCmd.MarkFlagRequired("project")
	evidentlyCmd.AddCommand(evidently_getExperimentCmd)
}
