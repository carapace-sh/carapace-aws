package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_startIdMappingJobCmd = &cobra.Command{
	Use:   "start-id-mapping-job",
	Short: "Starts the `IdMappingJob` of a workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_startIdMappingJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(entityresolution_startIdMappingJobCmd).Standalone()

		entityresolution_startIdMappingJobCmd.Flags().String("job-type", "", "The job type for the ID mapping job.")
		entityresolution_startIdMappingJobCmd.Flags().String("output-source-config", "", "A list of `OutputSource` objects.")
		entityresolution_startIdMappingJobCmd.Flags().String("workflow-name", "", "The name of the ID mapping job to be retrieved.")
		entityresolution_startIdMappingJobCmd.MarkFlagRequired("workflow-name")
	})
	entityresolutionCmd.AddCommand(entityresolution_startIdMappingJobCmd)
}
