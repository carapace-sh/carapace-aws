package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecr_createRepositoryCmd = &cobra.Command{
	Use:   "create-repository",
	Short: "Creates a repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecr_createRepositoryCmd).Standalone()

	ecr_createRepositoryCmd.Flags().String("encryption-configuration", "", "The encryption configuration for the repository.")
	ecr_createRepositoryCmd.Flags().String("image-scanning-configuration", "", "The image scanning configuration for the repository.")
	ecr_createRepositoryCmd.Flags().String("image-tag-mutability", "", "The tag mutability setting for the repository.")
	ecr_createRepositoryCmd.Flags().String("image-tag-mutability-exclusion-filters", "", "Creates a repository with a list of filters that define which image tags can override the default image tag mutability setting.")
	ecr_createRepositoryCmd.Flags().String("registry-id", "", "The Amazon Web Services account ID associated with the registry to create the repository.")
	ecr_createRepositoryCmd.Flags().String("repository-name", "", "The name to use for the repository.")
	ecr_createRepositoryCmd.Flags().String("tags", "", "The metadata that you apply to the repository to help you categorize and organize them.")
	ecr_createRepositoryCmd.MarkFlagRequired("repository-name")
	ecrCmd.AddCommand(ecr_createRepositoryCmd)
}
