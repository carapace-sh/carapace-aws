package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_createMatchingWorkflowCmd = &cobra.Command{
	Use:   "create-matching-workflow",
	Short: "Creates a matching workflow that defines the configuration for a data processing job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_createMatchingWorkflowCmd).Standalone()

	entityresolution_createMatchingWorkflowCmd.Flags().String("description", "", "A description of the workflow.")
	entityresolution_createMatchingWorkflowCmd.Flags().String("incremental-run-config", "", "Optional.")
	entityresolution_createMatchingWorkflowCmd.Flags().String("input-source-config", "", "A list of `InputSource` objects, which have the fields `InputSourceARN` and `SchemaName`.")
	entityresolution_createMatchingWorkflowCmd.Flags().String("output-source-config", "", "A list of `OutputSource` objects, each of which contains fields `outputS3Path`, `applyNormalization`, `KMSArn`, and `output`.")
	entityresolution_createMatchingWorkflowCmd.Flags().String("resolution-techniques", "", "An object which defines the `resolutionType` and the `ruleBasedProperties`.")
	entityresolution_createMatchingWorkflowCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM role.")
	entityresolution_createMatchingWorkflowCmd.Flags().String("tags", "", "The tags used to organize, track, or control access for this resource.")
	entityresolution_createMatchingWorkflowCmd.Flags().String("workflow-name", "", "The name of the workflow.")
	entityresolution_createMatchingWorkflowCmd.MarkFlagRequired("input-source-config")
	entityresolution_createMatchingWorkflowCmd.MarkFlagRequired("output-source-config")
	entityresolution_createMatchingWorkflowCmd.MarkFlagRequired("resolution-techniques")
	entityresolution_createMatchingWorkflowCmd.MarkFlagRequired("role-arn")
	entityresolution_createMatchingWorkflowCmd.MarkFlagRequired("workflow-name")
	entityresolutionCmd.AddCommand(entityresolution_createMatchingWorkflowCmd)
}
