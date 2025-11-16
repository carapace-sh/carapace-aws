package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_createDomainAssociationCmd = &cobra.Command{
	Use:   "create-domain-association",
	Short: "Creates a new domain association for an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_createDomainAssociationCmd).Standalone()

	amplify_createDomainAssociationCmd.Flags().String("app-id", "", "The unique ID for an Amplify app.")
	amplify_createDomainAssociationCmd.Flags().String("auto-sub-domain-creation-patterns", "", "Sets the branch patterns for automatic subdomain creation.")
	amplify_createDomainAssociationCmd.Flags().String("auto-sub-domain-iamrole", "", "The required AWS Identity and Access Management (IAM) service role for the Amazon Resource Name (ARN) for automatically creating subdomains.")
	amplify_createDomainAssociationCmd.Flags().String("certificate-settings", "", "The type of SSL/TLS certificate to use for your custom domain.")
	amplify_createDomainAssociationCmd.Flags().String("domain-name", "", "The domain name for the domain association.")
	amplify_createDomainAssociationCmd.Flags().String("enable-auto-sub-domain", "", "Enables the automated creation of subdomains for branches.")
	amplify_createDomainAssociationCmd.Flags().String("sub-domain-settings", "", "The setting for the subdomain.")
	amplify_createDomainAssociationCmd.MarkFlagRequired("app-id")
	amplify_createDomainAssociationCmd.MarkFlagRequired("domain-name")
	amplify_createDomainAssociationCmd.MarkFlagRequired("sub-domain-settings")
	amplifyCmd.AddCommand(amplify_createDomainAssociationCmd)
}
