package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_transferDomainToAnotherAwsAccountCmd = &cobra.Command{
	Use:   "transfer-domain-to-another-aws-account",
	Short: "Transfers a domain from the current Amazon Web Services account to another Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_transferDomainToAnotherAwsAccountCmd).Standalone()

	route53domains_transferDomainToAnotherAwsAccountCmd.Flags().String("account-id", "", "The account ID of the Amazon Web Services account that you want to transfer the domain to, for example, `111122223333`.")
	route53domains_transferDomainToAnotherAwsAccountCmd.Flags().String("domain-name", "", "The name of the domain that you want to transfer from the current Amazon Web Services account to another account.")
	route53domains_transferDomainToAnotherAwsAccountCmd.MarkFlagRequired("account-id")
	route53domains_transferDomainToAnotherAwsAccountCmd.MarkFlagRequired("domain-name")
	route53domainsCmd.AddCommand(route53domains_transferDomainToAnotherAwsAccountCmd)
}
