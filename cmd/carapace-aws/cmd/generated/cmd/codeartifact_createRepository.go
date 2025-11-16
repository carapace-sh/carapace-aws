package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_createRepositoryCmd = &cobra.Command{
	Use:   "create-repository",
	Short: "Creates a repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_createRepositoryCmd).Standalone()

	codeartifact_createRepositoryCmd.Flags().String("description", "", "A description of the created repository.")
	codeartifact_createRepositoryCmd.Flags().String("domain", "", "The name of the domain that contains the created repository.")
	codeartifact_createRepositoryCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
	codeartifact_createRepositoryCmd.Flags().String("repository", "", "The name of the repository to create.")
	codeartifact_createRepositoryCmd.Flags().String("tags", "", "One or more tag key-value pairs for the repository.")
	codeartifact_createRepositoryCmd.Flags().String("upstreams", "", "A list of upstream repositories to associate with the repository.")
	codeartifact_createRepositoryCmd.MarkFlagRequired("domain")
	codeartifact_createRepositoryCmd.MarkFlagRequired("repository")
	codeartifactCmd.AddCommand(codeartifact_createRepositoryCmd)
}
