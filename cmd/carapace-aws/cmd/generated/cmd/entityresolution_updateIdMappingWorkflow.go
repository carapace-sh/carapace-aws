package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_updateIdMappingWorkflowCmd = &cobra.Command{
	Use:   "update-id-mapping-workflow",
	Short: "Updates an existing `IdMappingWorkflow`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_updateIdMappingWorkflowCmd).Standalone()

	entityresolution_updateIdMappingWorkflowCmd.Flags().String("description", "", "A description of the workflow.")
	entityresolution_updateIdMappingWorkflowCmd.Flags().String("id-mapping-techniques", "", "An object which defines the ID mapping technique and any additional configurations.")
	entityresolution_updateIdMappingWorkflowCmd.Flags().String("incremental-run-config", "", "The incremental run configuration for the update ID mapping workflow.")
	entityresolution_updateIdMappingWorkflowCmd.Flags().String("input-source-config", "", "A list of `InputSource` objects, which have the fields `InputSourceARN` and `SchemaName`.")
	entityresolution_updateIdMappingWorkflowCmd.Flags().String("output-source-config", "", "A list of `OutputSource` objects, each of which contains fields `outputS3Path` and `KMSArn`.")
	entityresolution_updateIdMappingWorkflowCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM role.")
	entityresolution_updateIdMappingWorkflowCmd.Flags().String("workflow-name", "", "The name of the workflow.")
	entityresolution_updateIdMappingWorkflowCmd.MarkFlagRequired("id-mapping-techniques")
	entityresolution_updateIdMappingWorkflowCmd.MarkFlagRequired("input-source-config")
	entityresolution_updateIdMappingWorkflowCmd.MarkFlagRequired("workflow-name")
	entityresolutionCmd.AddCommand(entityresolution_updateIdMappingWorkflowCmd)
}
