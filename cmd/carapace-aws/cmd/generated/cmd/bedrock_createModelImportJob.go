package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_createModelImportJobCmd = &cobra.Command{
	Use:   "create-model-import-job",
	Short: "Creates a model import job to import model that you have customized in other environments, such as Amazon SageMaker.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_createModelImportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bedrock_createModelImportJobCmd).Standalone()

		bedrock_createModelImportJobCmd.Flags().String("client-request-token", "", "A unique, case-sensitive identifier to ensure that the API request completes no more than one time.")
		bedrock_createModelImportJobCmd.Flags().String("imported-model-kms-key-id", "", "The imported model is encrypted at rest using this key.")
		bedrock_createModelImportJobCmd.Flags().String("imported-model-name", "", "The name of the imported model.")
		bedrock_createModelImportJobCmd.Flags().String("imported-model-tags", "", "Tags to attach to the imported model.")
		bedrock_createModelImportJobCmd.Flags().String("job-name", "", "The name of the import job.")
		bedrock_createModelImportJobCmd.Flags().String("job-tags", "", "Tags to attach to this import job.")
		bedrock_createModelImportJobCmd.Flags().String("model-data-source", "", "The data source for the imported model.")
		bedrock_createModelImportJobCmd.Flags().String("role-arn", "", "The Amazon Resource Name (ARN) of the model import job.")
		bedrock_createModelImportJobCmd.Flags().String("vpc-config", "", "VPC configuration parameters for the private Virtual Private Cloud (VPC) that contains the resources you are using for the import job.")
		bedrock_createModelImportJobCmd.MarkFlagRequired("imported-model-name")
		bedrock_createModelImportJobCmd.MarkFlagRequired("job-name")
		bedrock_createModelImportJobCmd.MarkFlagRequired("model-data-source")
		bedrock_createModelImportJobCmd.MarkFlagRequired("role-arn")
	})
	bedrockCmd.AddCommand(bedrock_createModelImportJobCmd)
}
