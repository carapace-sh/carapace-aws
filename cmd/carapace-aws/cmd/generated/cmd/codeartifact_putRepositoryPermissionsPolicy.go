package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_putRepositoryPermissionsPolicyCmd = &cobra.Command{
	Use:   "put-repository-permissions-policy",
	Short: "Sets the resource policy on a repository that specifies permissions to access it.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_putRepositoryPermissionsPolicyCmd).Standalone()

	codeartifact_putRepositoryPermissionsPolicyCmd.Flags().String("domain", "", "The name of the domain containing the repository to set the resource policy on.")
	codeartifact_putRepositoryPermissionsPolicyCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
	codeartifact_putRepositoryPermissionsPolicyCmd.Flags().String("policy-document", "", "A valid displayable JSON Aspen policy string to be set as the access control resource policy on the provided repository.")
	codeartifact_putRepositoryPermissionsPolicyCmd.Flags().String("policy-revision", "", "Sets the revision of the resource policy that specifies permissions to access the repository.")
	codeartifact_putRepositoryPermissionsPolicyCmd.Flags().String("repository", "", "The name of the repository to set the resource policy on.")
	codeartifact_putRepositoryPermissionsPolicyCmd.MarkFlagRequired("domain")
	codeartifact_putRepositoryPermissionsPolicyCmd.MarkFlagRequired("policy-document")
	codeartifact_putRepositoryPermissionsPolicyCmd.MarkFlagRequired("repository")
	codeartifactCmd.AddCommand(codeartifact_putRepositoryPermissionsPolicyCmd)
}
