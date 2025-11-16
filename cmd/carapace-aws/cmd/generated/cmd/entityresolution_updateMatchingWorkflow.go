package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var entityresolution_updateMatchingWorkflowCmd = &cobra.Command{
	Use:   "update-matching-workflow",
	Short: "Updates an existing matching workflow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(entityresolution_updateMatchingWorkflowCmd).Standalone()

	entityresolution_updateMatchingWorkflowCmd.Flags().String("description", "", "A description of the workflow.")
	entityresolution_updateMatchingWorkflowCmd.Flags().String("incremental-run-config", "", "Optional.")
	entityresolution_updateMatchingWorkflowCmd.Flags().String("input-source-config", "", "A list of `InputSource` objects, which have the fields `InputSourceARN` and `SchemaName`.")
	entityresolution_updateMatchingWorkflowCmd.Flags().String("output-source-config", "", "A list of `OutputSource` objects, each of which contains fields `outputS3Path`, `applyNormalization`, `KMSArn`, and `output`.")
	entityresolution_updateMatchingWorkflowCmd.Flags().String("resolution-techniques", "", "An object which defines the `resolutionType` and the `ruleBasedProperties`.")
	entityresolution_updateMatchingWorkflowCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the IAM role.")
	entityresolution_updateMatchingWorkflowCmd.Flags().String("workflow-name", "", "The name of the workflow to be retrieved.")
	entityresolution_updateMatchingWorkflowCmd.MarkFlagRequired("input-source-config")
	entityresolution_updateMatchingWorkflowCmd.MarkFlagRequired("output-source-config")
	entityresolution_updateMatchingWorkflowCmd.MarkFlagRequired("resolution-techniques")
	entityresolution_updateMatchingWorkflowCmd.MarkFlagRequired("role-arn")
	entityresolution_updateMatchingWorkflowCmd.MarkFlagRequired("workflow-name")
	entityresolutionCmd.AddCommand(entityresolution_updateMatchingWorkflowCmd)
}
