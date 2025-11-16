package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_updateRepositoryCmd = &cobra.Command{
	Use:   "update-repository",
	Short: "Update the properties of a repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_updateRepositoryCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeartifact_updateRepositoryCmd).Standalone()

		codeartifact_updateRepositoryCmd.Flags().String("description", "", "An updated repository description.")
		codeartifact_updateRepositoryCmd.Flags().String("domain", "", "The name of the domain associated with the repository to update.")
		codeartifact_updateRepositoryCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
		codeartifact_updateRepositoryCmd.Flags().String("repository", "", "The name of the repository to update.")
		codeartifact_updateRepositoryCmd.Flags().String("upstreams", "", "A list of upstream repositories to associate with the repository.")
		codeartifact_updateRepositoryCmd.MarkFlagRequired("domain")
		codeartifact_updateRepositoryCmd.MarkFlagRequired("repository")
	})
	codeartifactCmd.AddCommand(codeartifact_updateRepositoryCmd)
}
