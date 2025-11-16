package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_startMatchingJobCmd = &cobra.Command{
	Use:   "start-matching-job",
	Short: "Starts the `MatchingJob` of a workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_startMatchingJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(entityresolution_startMatchingJobCmd).Standalone()

		entityresolution_startMatchingJobCmd.Flags().String("workflow-name", "", "The name of the matching job to be retrieved.")
		entityresolution_startMatchingJobCmd.MarkFlagRequired("workflow-name")
	})
	entityresolutionCmd.AddCommand(entityresolution_startMatchingJobCmd)
}
