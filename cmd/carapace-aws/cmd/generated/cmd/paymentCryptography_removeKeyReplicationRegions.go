package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptography_removeKeyReplicationRegionsCmd = &cobra.Command{
	Use:   "remove-key-replication-regions",
	Short: "Removes Replication Regions from an existing Amazon Web Services Payment Cryptography key, disabling the key's availability for cryptographic operations in the specified Amazon Web Services Regions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptography_removeKeyReplicationRegionsCmd).Standalone()

	paymentCryptography_removeKeyReplicationRegionsCmd.Flags().String("key-identifier", "", "The key identifier (ARN or alias) of the key from which to remove replication regions.")
	paymentCryptography_removeKeyReplicationRegionsCmd.Flags().String("replication-regions", "", "The list of Amazon Web Services Regions to remove from the key's replication configuration.")
	paymentCryptography_removeKeyReplicationRegionsCmd.MarkFlagRequired("key-identifier")
	paymentCryptography_removeKeyReplicationRegionsCmd.MarkFlagRequired("replication-regions")
	paymentCryptographyCmd.AddCommand(paymentCryptography_removeKeyReplicationRegionsCmd)
}
