package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_deleteRepositoryCmd = &cobra.Command{
	Use:   "delete-repository",
	Short: "Deletes a repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_deleteRepositoryCmd).Standalone()

	codeartifact_deleteRepositoryCmd.Flags().String("domain", "", "The name of the domain that contains the repository to delete.")
	codeartifact_deleteRepositoryCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
	codeartifact_deleteRepositoryCmd.Flags().String("repository", "", "The name of the repository to delete.")
	codeartifact_deleteRepositoryCmd.MarkFlagRequired("domain")
	codeartifact_deleteRepositoryCmd.MarkFlagRequired("repository")
	codeartifactCmd.AddCommand(codeartifact_deleteRepositoryCmd)
}
