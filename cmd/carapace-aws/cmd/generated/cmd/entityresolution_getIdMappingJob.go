package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_getIdMappingJobCmd = &cobra.Command{
	Use:   "get-id-mapping-job",
	Short: "Returns the status, metrics, and errors (if there are any) that are associated with a job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_getIdMappingJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(entityresolution_getIdMappingJobCmd).Standalone()

		entityresolution_getIdMappingJobCmd.Flags().String("job-id", "", "The ID of the job.")
		entityresolution_getIdMappingJobCmd.Flags().String("workflow-name", "", "The name of the workflow.")
		entityresolution_getIdMappingJobCmd.MarkFlagRequired("job-id")
		entityresolution_getIdMappingJobCmd.MarkFlagRequired("workflow-name")
	})
	entityresolutionCmd.AddCommand(entityresolution_getIdMappingJobCmd)
}
