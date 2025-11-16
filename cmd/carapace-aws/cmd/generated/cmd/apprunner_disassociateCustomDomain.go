package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apprunner_disassociateCustomDomainCmd = &cobra.Command{
	Use:   "disassociate-custom-domain",
	Short: "Disassociate a custom domain name from an App Runner service.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apprunner_disassociateCustomDomainCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apprunner_disassociateCustomDomainCmd).Standalone()

		apprunner_disassociateCustomDomainCmd.Flags().String("domain-name", "", "The domain name that you want to disassociate from the App Runner service.")
		apprunner_disassociateCustomDomainCmd.Flags().String("service-arn", "", "The Amazon Resource Name (ARN) of the App Runner service that you want to disassociate a custom domain name from.")
		apprunner_disassociateCustomDomainCmd.MarkFlagRequired("domain-name")
		apprunner_disassociateCustomDomainCmd.MarkFlagRequired("service-arn")
	})
	apprunnerCmd.AddCommand(apprunner_disassociateCustomDomainCmd)
}
