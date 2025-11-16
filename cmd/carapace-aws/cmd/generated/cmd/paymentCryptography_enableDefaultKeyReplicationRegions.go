package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptography_enableDefaultKeyReplicationRegionsCmd = &cobra.Command{
	Use:   "enable-default-key-replication-regions",
	Short: "Enables [Multi-Region key replication](https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-multi-region-replication.html) settings for your Amazon Web Services account, causing new keys to be automatically replicated to the specified Amazon Web Services Regions when created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptography_enableDefaultKeyReplicationRegionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(paymentCryptography_enableDefaultKeyReplicationRegionsCmd).Standalone()

		paymentCryptography_enableDefaultKeyReplicationRegionsCmd.Flags().String("replication-regions", "", "The list of Amazon Web Services Regions to enable as default replication regions for the Amazon Web Services account for [Multi-Region key replication](https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-multi-region-replication.html).")
		paymentCryptography_enableDefaultKeyReplicationRegionsCmd.MarkFlagRequired("replication-regions")
	})
	paymentCryptographyCmd.AddCommand(paymentCryptography_enableDefaultKeyReplicationRegionsCmd)
}
