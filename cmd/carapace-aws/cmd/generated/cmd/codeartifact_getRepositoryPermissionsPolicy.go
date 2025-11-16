package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_getRepositoryPermissionsPolicyCmd = &cobra.Command{
	Use:   "get-repository-permissions-policy",
	Short: "Returns the resource policy that is set on a repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_getRepositoryPermissionsPolicyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codeartifact_getRepositoryPermissionsPolicyCmd).Standalone()

		codeartifact_getRepositoryPermissionsPolicyCmd.Flags().String("domain", "", "The name of the domain containing the repository whose associated resource policy is to be retrieved.")
		codeartifact_getRepositoryPermissionsPolicyCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
		codeartifact_getRepositoryPermissionsPolicyCmd.Flags().String("repository", "", "The name of the repository whose associated resource policy is to be retrieved.")
		codeartifact_getRepositoryPermissionsPolicyCmd.MarkFlagRequired("domain")
		codeartifact_getRepositoryPermissionsPolicyCmd.MarkFlagRequired("repository")
	})
	codeartifactCmd.AddCommand(codeartifact_getRepositoryPermissionsPolicyCmd)
}
