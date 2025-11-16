package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptography_disableDefaultKeyReplicationRegionsCmd = &cobra.Command{
	Use:   "disable-default-key-replication-regions",
	Short: "Disables [Multi-Region key replication](https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-multi-region-replication.html) settings for the specified Amazon Web Services Regions in your Amazon Web Services account, preventing new keys from being automatically replicated to those regions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptography_disableDefaultKeyReplicationRegionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(paymentCryptography_disableDefaultKeyReplicationRegionsCmd).Standalone()

		paymentCryptography_disableDefaultKeyReplicationRegionsCmd.Flags().String("replication-regions", "", "The list of Amazon Web Services Regions to remove from the account's default replication regions.")
		paymentCryptography_disableDefaultKeyReplicationRegionsCmd.MarkFlagRequired("replication-regions")
	})
	paymentCryptographyCmd.AddCommand(paymentCryptography_disableDefaultKeyReplicationRegionsCmd)
}
