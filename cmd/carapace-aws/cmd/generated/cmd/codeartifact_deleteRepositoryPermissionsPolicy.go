package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_deleteRepositoryPermissionsPolicyCmd = &cobra.Command{
	Use:   "delete-repository-permissions-policy",
	Short: "Deletes the resource policy that is set on a repository.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_deleteRepositoryPermissionsPolicyCmd).Standalone()

	codeartifact_deleteRepositoryPermissionsPolicyCmd.Flags().String("domain", "", "The name of the domain that contains the repository associated with the resource policy to be deleted.")
	codeartifact_deleteRepositoryPermissionsPolicyCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
	codeartifact_deleteRepositoryPermissionsPolicyCmd.Flags().String("policy-revision", "", "The revision of the repository's resource policy to be deleted.")
	codeartifact_deleteRepositoryPermissionsPolicyCmd.Flags().String("repository", "", "The name of the repository that is associated with the resource policy to be deleted")
	codeartifact_deleteRepositoryPermissionsPolicyCmd.MarkFlagRequired("domain")
	codeartifact_deleteRepositoryPermissionsPolicyCmd.MarkFlagRequired("repository")
	codeartifactCmd.AddCommand(codeartifact_deleteRepositoryPermissionsPolicyCmd)
}
