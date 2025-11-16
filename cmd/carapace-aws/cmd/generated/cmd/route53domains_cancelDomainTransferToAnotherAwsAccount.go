package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_cancelDomainTransferToAnotherAwsAccountCmd = &cobra.Command{
	Use:   "cancel-domain-transfer-to-another-aws-account",
	Short: "Cancels the transfer of a domain from the current Amazon Web Services account to another Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_cancelDomainTransferToAnotherAwsAccountCmd).Standalone()

	route53domains_cancelDomainTransferToAnotherAwsAccountCmd.Flags().String("domain-name", "", "The name of the domain for which you want to cancel the transfer to another Amazon Web Services account.")
	route53domains_cancelDomainTransferToAnotherAwsAccountCmd.MarkFlagRequired("domain-name")
	route53domainsCmd.AddCommand(route53domains_cancelDomainTransferToAnotherAwsAccountCmd)
}
