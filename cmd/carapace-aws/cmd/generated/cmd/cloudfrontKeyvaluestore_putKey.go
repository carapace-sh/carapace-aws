package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfrontKeyvaluestore_putKeyCmd = &cobra.Command{
	Use:   "put-key",
	Short: "Creates a new key value pair or replaces the value of an existing key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfrontKeyvaluestore_putKeyCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfrontKeyvaluestore_putKeyCmd).Standalone()

		cloudfrontKeyvaluestore_putKeyCmd.Flags().String("if-match", "", "The current version (ETag) of the Key Value Store that you are putting keys into, which you can get using DescribeKeyValueStore.")
		cloudfrontKeyvaluestore_putKeyCmd.Flags().String("key", "", "The key to put.")
		cloudfrontKeyvaluestore_putKeyCmd.Flags().String("kvs-arn", "", "The Amazon Resource Name (ARN) of the Key Value Store.")
		cloudfrontKeyvaluestore_putKeyCmd.Flags().String("value", "", "The value to put.")
		cloudfrontKeyvaluestore_putKeyCmd.MarkFlagRequired("if-match")
		cloudfrontKeyvaluestore_putKeyCmd.MarkFlagRequired("key")
		cloudfrontKeyvaluestore_putKeyCmd.MarkFlagRequired("kvs-arn")
		cloudfrontKeyvaluestore_putKeyCmd.MarkFlagRequired("value")
	})
	cloudfrontKeyvaluestoreCmd.AddCommand(cloudfrontKeyvaluestore_putKeyCmd)
}
