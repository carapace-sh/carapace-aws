package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_associateCustomDomainCmd = &cobra.Command{
	Use:   "associate-custom-domain",
	Short: "Associate your own domain name with the App Runner subdomain URL of your App Runner service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_associateCustomDomainCmd).Standalone()

	apprunner_associateCustomDomainCmd.Flags().String("domain-name", "", "A custom domain endpoint to associate.")
	apprunner_associateCustomDomainCmd.Flags().String("enable-wwwsubdomain", "", "Set to `true` to associate the subdomain `www.DomainName` with the App Runner service in addition to the base domain.")
	apprunner_associateCustomDomainCmd.Flags().String("service-arn", "", "The Amazon Resource Name (ARN) of the App Runner service that you want to associate a custom domain name with.")
	apprunner_associateCustomDomainCmd.MarkFlagRequired("domain-name")
	apprunner_associateCustomDomainCmd.MarkFlagRequired("service-arn")
	apprunnerCmd.AddCommand(apprunner_associateCustomDomainCmd)
}
