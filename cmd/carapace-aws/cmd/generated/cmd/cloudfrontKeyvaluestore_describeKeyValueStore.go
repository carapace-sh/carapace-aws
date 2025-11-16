package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfrontKeyvaluestore_describeKeyValueStoreCmd = &cobra.Command{
	Use:   "describe-key-value-store",
	Short: "Returns metadata information about Key Value Store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfrontKeyvaluestore_describeKeyValueStoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfrontKeyvaluestore_describeKeyValueStoreCmd).Standalone()

		cloudfrontKeyvaluestore_describeKeyValueStoreCmd.Flags().String("kvs-arn", "", "The Amazon Resource Name (ARN) of the Key Value Store.")
		cloudfrontKeyvaluestore_describeKeyValueStoreCmd.MarkFlagRequired("kvs-arn")
	})
	cloudfrontKeyvaluestoreCmd.AddCommand(cloudfrontKeyvaluestore_describeKeyValueStoreCmd)
}
