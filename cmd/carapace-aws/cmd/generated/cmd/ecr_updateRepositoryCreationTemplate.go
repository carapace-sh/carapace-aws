package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_updateRepositoryCreationTemplateCmd = &cobra.Command{
	Use:   "update-repository-creation-template",
	Short: "Updates an existing repository creation template.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_updateRepositoryCreationTemplateCmd).Standalone()

	ecr_updateRepositoryCreationTemplateCmd.Flags().String("applied-for", "", "Updates the list of enumerable strings representing the Amazon ECR repository creation scenarios that this template will apply towards.")
	ecr_updateRepositoryCreationTemplateCmd.Flags().String("custom-role-arn", "", "The ARN of the role to be assumed by Amazon ECR.")
	ecr_updateRepositoryCreationTemplateCmd.Flags().String("description", "", "A description for the repository creation template.")
	ecr_updateRepositoryCreationTemplateCmd.Flags().String("encryption-configuration", "", "")
	ecr_updateRepositoryCreationTemplateCmd.Flags().String("image-tag-mutability", "", "Updates the tag mutability setting for the repository.")
	ecr_updateRepositoryCreationTemplateCmd.Flags().String("image-tag-mutability-exclusion-filters", "", "Updates a repository with filters that define which image tags can override the default image tag mutability setting.")
	ecr_updateRepositoryCreationTemplateCmd.Flags().String("lifecycle-policy", "", "Updates the lifecycle policy associated with the specified repository creation template.")
	ecr_updateRepositoryCreationTemplateCmd.Flags().String("prefix", "", "The repository namespace prefix that matches an existing repository creation template in the registry.")
	ecr_updateRepositoryCreationTemplateCmd.Flags().String("repository-policy", "", "Updates the repository policy created using the template.")
	ecr_updateRepositoryCreationTemplateCmd.Flags().String("resource-tags", "", "The metadata to apply to the repository to help you categorize and organize.")
	ecr_updateRepositoryCreationTemplateCmd.MarkFlagRequired("prefix")
	ecrCmd.AddCommand(ecr_updateRepositoryCreationTemplateCmd)
}
