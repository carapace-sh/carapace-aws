package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_createRepositoryCreationTemplateCmd = &cobra.Command{
	Use:   "create-repository-creation-template",
	Short: "Creates a repository creation template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_createRepositoryCreationTemplateCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(ecr_createRepositoryCreationTemplateCmd).Standalone()

		ecr_createRepositoryCreationTemplateCmd.Flags().String("applied-for", "", "A list of enumerable strings representing the Amazon ECR repository creation scenarios that this template will apply towards.")
		ecr_createRepositoryCreationTemplateCmd.Flags().String("custom-role-arn", "", "The ARN of the role to be assumed by Amazon ECR.")
		ecr_createRepositoryCreationTemplateCmd.Flags().String("description", "", "A description for the repository creation template.")
		ecr_createRepositoryCreationTemplateCmd.Flags().String("encryption-configuration", "", "The encryption configuration to use for repositories created using the template.")
		ecr_createRepositoryCreationTemplateCmd.Flags().String("image-tag-mutability", "", "The tag mutability setting for the repository.")
		ecr_createRepositoryCreationTemplateCmd.Flags().String("image-tag-mutability-exclusion-filters", "", "Creates a repository creation template with a list of filters that define which image tags can override the default image tag mutability setting.")
		ecr_createRepositoryCreationTemplateCmd.Flags().String("lifecycle-policy", "", "The lifecycle policy to use for repositories created using the template.")
		ecr_createRepositoryCreationTemplateCmd.Flags().String("prefix", "", "The repository namespace prefix to associate with the template.")
		ecr_createRepositoryCreationTemplateCmd.Flags().String("repository-policy", "", "The repository policy to apply to repositories created using the template.")
		ecr_createRepositoryCreationTemplateCmd.Flags().String("resource-tags", "", "The metadata to apply to the repository to help you categorize and organize.")
		ecr_createRepositoryCreationTemplateCmd.MarkFlagRequired("applied-for")
		ecr_createRepositoryCreationTemplateCmd.MarkFlagRequired("prefix")
	})
	ecrCmd.AddCommand(ecr_createRepositoryCreationTemplateCmd)
}
