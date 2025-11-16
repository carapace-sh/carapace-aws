package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_deleteDomainPermissionsPolicyCmd = &cobra.Command{
	Use:   "delete-domain-permissions-policy",
	Short: "Deletes the resource policy set on a domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_deleteDomainPermissionsPolicyCmd).Standalone()

	codeartifact_deleteDomainPermissionsPolicyCmd.Flags().String("domain", "", "The name of the domain associated with the resource policy to be deleted.")
	codeartifact_deleteDomainPermissionsPolicyCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
	codeartifact_deleteDomainPermissionsPolicyCmd.Flags().String("policy-revision", "", "The current revision of the resource policy to be deleted.")
	codeartifact_deleteDomainPermissionsPolicyCmd.MarkFlagRequired("domain")
	codeartifactCmd.AddCommand(codeartifact_deleteDomainPermissionsPolicyCmd)
}
