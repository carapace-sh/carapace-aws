package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codeartifact_putDomainPermissionsPolicyCmd = &cobra.Command{
	Use:   "put-domain-permissions-policy",
	Short: "Sets a resource policy on a domain that specifies permissions to access it.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codeartifact_putDomainPermissionsPolicyCmd).Standalone()

	codeartifact_putDomainPermissionsPolicyCmd.Flags().String("domain", "", "The name of the domain on which to set the resource policy.")
	codeartifact_putDomainPermissionsPolicyCmd.Flags().String("domain-owner", "", "The 12-digit account number of the Amazon Web Services account that owns the domain.")
	codeartifact_putDomainPermissionsPolicyCmd.Flags().String("policy-document", "", "A valid displayable JSON Aspen policy string to be set as the access control resource policy on the provided domain.")
	codeartifact_putDomainPermissionsPolicyCmd.Flags().String("policy-revision", "", "The current revision of the resource policy to be set.")
	codeartifact_putDomainPermissionsPolicyCmd.MarkFlagRequired("domain")
	codeartifact_putDomainPermissionsPolicyCmd.MarkFlagRequired("policy-document")
	codeartifactCmd.AddCommand(codeartifact_putDomainPermissionsPolicyCmd)
}
