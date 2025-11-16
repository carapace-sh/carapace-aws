package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptography_addKeyReplicationRegionsCmd = &cobra.Command{
	Use:   "add-key-replication-regions",
	Short: "Adds replication Amazon Web Services Regions to an existing Amazon Web Services Payment Cryptography key, enabling the key to be used for cryptographic operations in additional Amazon Web Services Regions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptography_addKeyReplicationRegionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(paymentCryptography_addKeyReplicationRegionsCmd).Standalone()

		paymentCryptography_addKeyReplicationRegionsCmd.Flags().String("key-identifier", "", "The key identifier (ARN or alias) of the key for which to add replication regions.")
		paymentCryptography_addKeyReplicationRegionsCmd.Flags().String("replication-regions", "", "The list of Amazon Web Services Regions to add to the key's replication configuration.")
		paymentCryptography_addKeyReplicationRegionsCmd.MarkFlagRequired("key-identifier")
		paymentCryptography_addKeyReplicationRegionsCmd.MarkFlagRequired("replication-regions")
	})
	paymentCryptographyCmd.AddCommand(paymentCryptography_addKeyReplicationRegionsCmd)
}
