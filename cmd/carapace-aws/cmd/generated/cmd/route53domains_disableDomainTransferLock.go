package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53domains_disableDomainTransferLockCmd = &cobra.Command{
	Use:   "disable-domain-transfer-lock",
	Short: "This operation removes the transfer lock on the domain (specifically the `clientTransferProhibited` status) to allow domain transfers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53domains_disableDomainTransferLockCmd).Standalone()

	route53domains_disableDomainTransferLockCmd.Flags().String("domain-name", "", "The name of the domain that you want to remove the transfer lock for.")
	route53domains_disableDomainTransferLockCmd.MarkFlagRequired("domain-name")
	route53domainsCmd.AddCommand(route53domains_disableDomainTransferLockCmd)
}
