package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var paymentCryptography_getDefaultKeyReplicationRegionsCmd = &cobra.Command{
	Use:   "get-default-key-replication-regions",
	Short: "Retrieves the list of Amazon Web Services Regions where [Multi-Region key replication](https://docs.aws.amazon.com/payment-cryptography/latest/userguide/keys-multi-region-replication.html) is currently enabled for your Amazon Web Services account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(paymentCryptography_getDefaultKeyReplicationRegionsCmd).Standalone()

	paymentCryptographyCmd.AddCommand(paymentCryptography_getDefaultKeyReplicationRegionsCmd)
}
