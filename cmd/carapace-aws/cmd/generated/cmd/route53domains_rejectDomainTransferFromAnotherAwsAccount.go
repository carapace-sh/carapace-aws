package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_rejectDomainTransferFromAnotherAwsAccountCmd = &cobra.Command{
	Use:   "reject-domain-transfer-from-another-aws-account",
	Short: "Rejects the transfer of a domain from another Amazon Web Services account to the current Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_rejectDomainTransferFromAnotherAwsAccountCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53domains_rejectDomainTransferFromAnotherAwsAccountCmd).Standalone()

		route53domains_rejectDomainTransferFromAnotherAwsAccountCmd.Flags().String("domain-name", "", "The name of the domain that was specified when another Amazon Web Services account submitted a [TransferDomainToAnotherAwsAccount](https://docs.aws.amazon.com/Route53/latest/APIReference/API_domains_TransferDomainToAnotherAwsAccount.html) request.")
		route53domains_rejectDomainTransferFromAnotherAwsAccountCmd.MarkFlagRequired("domain-name")
	})
	route53domainsCmd.AddCommand(route53domains_rejectDomainTransferFromAnotherAwsAccountCmd)
}
