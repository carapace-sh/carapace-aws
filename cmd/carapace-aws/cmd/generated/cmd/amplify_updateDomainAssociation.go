package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amplify_updateDomainAssociationCmd = &cobra.Command{
	Use:   "update-domain-association",
	Short: "Creates a new domain association for an Amplify app.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amplify_updateDomainAssociationCmd).Standalone()

	amplify_updateDomainAssociationCmd.Flags().String("app-id", "", "The unique ID for an Amplify app.")
	amplify_updateDomainAssociationCmd.Flags().String("auto-sub-domain-creation-patterns", "", "Sets the branch patterns for automatic subdomain creation.")
	amplify_updateDomainAssociationCmd.Flags().String("auto-sub-domain-iamrole", "", "The required AWS Identity and Access Management (IAM) service role for the Amazon Resource Name (ARN) for automatically creating subdomains.")
	amplify_updateDomainAssociationCmd.Flags().String("certificate-settings", "", "The type of SSL/TLS certificate to use for your custom domain.")
	amplify_updateDomainAssociationCmd.Flags().String("domain-name", "", "The name of the domain.")
	amplify_updateDomainAssociationCmd.Flags().String("enable-auto-sub-domain", "", "Enables the automated creation of subdomains for branches.")
	amplify_updateDomainAssociationCmd.Flags().String("sub-domain-settings", "", "Describes the settings for the subdomain.")
	amplify_updateDomainAssociationCmd.MarkFlagRequired("app-id")
	amplify_updateDomainAssociationCmd.MarkFlagRequired("domain-name")
	amplifyCmd.AddCommand(amplify_updateDomainAssociationCmd)
}
