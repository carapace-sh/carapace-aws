package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_getMatchingJobCmd = &cobra.Command{
	Use:   "get-matching-job",
	Short: "Returns the status, metrics, and errors (if there are any) that are associated with a job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_getMatchingJobCmd).Standalone()

	entityresolution_getMatchingJobCmd.Flags().String("job-id", "", "The ID of the job.")
	entityresolution_getMatchingJobCmd.Flags().String("workflow-name", "", "The name of the workflow.")
	entityresolution_getMatchingJobCmd.MarkFlagRequired("job-id")
	entityresolution_getMatchingJobCmd.MarkFlagRequired("workflow-name")
	entityresolutionCmd.AddCommand(entityresolution_getMatchingJobCmd)
}
