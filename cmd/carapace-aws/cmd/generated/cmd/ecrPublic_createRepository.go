package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ecrPublic_createRepositoryCmd = &cobra.Command{
	Use:   "create-repository",
	Short: "Creates a repository in a public registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ecrPublic_createRepositoryCmd).Standalone()

	ecrPublic_createRepositoryCmd.Flags().String("catalog-data", "", "The details about the repository that are publicly visible in the Amazon ECR Public Gallery.")
	ecrPublic_createRepositoryCmd.Flags().String("repository-name", "", "The name to use for the repository.")
	ecrPublic_createRepositoryCmd.Flags().String("tags", "", "The metadata that you apply to each repository to help categorize and organize your repositories.")
	ecrPublic_createRepositoryCmd.MarkFlagRequired("repository-name")
	ecrPublicCmd.AddCommand(ecrPublic_createRepositoryCmd)
}
