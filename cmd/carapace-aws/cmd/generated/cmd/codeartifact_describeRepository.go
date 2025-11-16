package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_describeRepositoryCmd = &cobra.Command{
	Use:   "describe-repository",
	Short: "Returns a `RepositoryDescription` object that contains detailed information about the requested repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_describeRepositoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeartifact_describeRepositoryCmd).Standalone()

		codeartifact_describeRepositoryCmd.Flags().String("domain", "", "The name of the domain that contains the repository to describe.")
		codeartifact_describeRepositoryCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
		codeartifact_describeRepositoryCmd.Flags().String("repository", "", "A string that specifies the name of the requested repository.")
		codeartifact_describeRepositoryCmd.MarkFlagRequired("domain")
		codeartifact_describeRepositoryCmd.MarkFlagRequired("repository")
	})
	codeartifactCmd.AddCommand(codeartifact_describeRepositoryCmd)
}
