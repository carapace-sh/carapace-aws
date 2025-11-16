package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_enableDomainTransferLockCmd = &cobra.Command{
	Use:   "enable-domain-transfer-lock",
	Short: "This operation sets the transfer lock on the domain (specifically the `clientTransferProhibited` status) to prevent domain transfers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_enableDomainTransferLockCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53domains_enableDomainTransferLockCmd).Standalone()

		route53domains_enableDomainTransferLockCmd.Flags().String("domain-name", "", "The name of the domain that you want to set the transfer lock for.")
		route53domains_enableDomainTransferLockCmd.MarkFlagRequired("domain-name")
	})
	route53domainsCmd.AddCommand(route53domains_enableDomainTransferLockCmd)
}
