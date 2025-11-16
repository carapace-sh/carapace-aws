package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_acceptDomainTransferFromAnotherAwsAccountCmd = &cobra.Command{
	Use:   "accept-domain-transfer-from-another-aws-account",
	Short: "Accepts the transfer of a domain from another Amazon Web Services account to the currentAmazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_acceptDomainTransferFromAnotherAwsAccountCmd).Standalone()

	route53domains_acceptDomainTransferFromAnotherAwsAccountCmd.Flags().String("domain-name", "", "The name of the domain that was specified when another Amazon Web Services account submitted a [TransferDomainToAnotherAwsAccount](https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_TransferDomainToAnotherAwsAccount.html) request.")
	route53domains_acceptDomainTransferFromAnotherAwsAccountCmd.Flags().String("password", "", "The password that was returned by the [TransferDomainToAnotherAwsAccount](https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_TransferDomainToAnotherAwsAccount.html) request.")
	route53domains_acceptDomainTransferFromAnotherAwsAccountCmd.MarkFlagRequired("domain-name")
	route53domains_acceptDomainTransferFromAnotherAwsAccountCmd.MarkFlagRequired("password")
	route53domainsCmd.AddCommand(route53domains_acceptDomainTransferFromAnotherAwsAccountCmd)
}
