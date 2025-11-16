package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_getDomainPermissionsPolicyCmd = &cobra.Command{
	Use:   "get-domain-permissions-policy",
	Short: "Returns the resource policy attached to the specified domain.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_getDomainPermissionsPolicyCmd).Standalone()

	codeartifact_getDomainPermissionsPolicyCmd.Flags().String("domain", "", "The name of the domain to which the resource policy is attached.")
	codeartifact_getDomainPermissionsPolicyCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
	codeartifact_getDomainPermissionsPolicyCmd.MarkFlagRequired("domain")
	codeartifactCmd.AddCommand(codeartifact_getDomainPermissionsPolicyCmd)
}
