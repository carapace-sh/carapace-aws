package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_createIdMappingWorkflowCmd = &cobra.Command{
	Use:   "create-id-mapping-workflow",
	Short: "Creates an `IdMappingWorkflow` object which stores the configuration of the data processing job to be run.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_createIdMappingWorkflowCmd).Standalone()

	entityresolution_createIdMappingWorkflowCmd.Flags().String("description", "", "A description of the workflow.")
	entityresolution_createIdMappingWorkflowCmd.Flags().String("id-mapping-techniques", "", "An object which defines the ID mapping technique and any additional configurations.")
	entityresolution_createIdMappingWorkflowCmd.Flags().String("incremental-run-config", "", "The incremental run configuration for the ID mapping workflow.")
	entityresolution_createIdMappingWorkflowCmd.Flags().String("input-source-config", "", "A list of `InputSource` objects, which have the fields `InputSourceARN` and `SchemaName`.")
	entityresolution_createIdMappingWorkflowCmd.Flags().String("output-source-config", "", "A list of `IdMappingWorkflowOutputSource` objects, each of which contains fields `outputS3Path` and `KMSArn`.")
	entityresolution_createIdMappingWorkflowCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM role.")
	entityresolution_createIdMappingWorkflowCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
	entityresolution_createIdMappingWorkflowCmd.Flags().String("workflow-name", "", "The name of the workflow.")
	entityresolution_createIdMappingWorkflowCmd.MarkFlagRequired("id-mapping-techniques")
	entityresolution_createIdMappingWorkflowCmd.MarkFlagRequired("input-source-config")
	entityresolution_createIdMappingWorkflowCmd.MarkFlagRequired("workflow-name")
	entityresolutionCmd.AddCommand(entityresolution_createIdMappingWorkflowCmd)
}
