package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfrontKeyvaluestore_getKeyCmd = &cobra.Command{
	Use:   "get-key",
	Short: "Returns a key value pair.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfrontKeyvaluestore_getKeyCmd).Standalone()

	cloudfrontKeyvaluestore_getKeyCmd.Flags().String("key", "", "The key to get.")
	cloudfrontKeyvaluestore_getKeyCmd.Flags().String("kvs-arn", "", "The Amazon Resource Name (ARN) of the Key Value Store.")
	cloudfrontKeyvaluestore_getKeyCmd.MarkFlagRequired("key")
	cloudfrontKeyvaluestore_getKeyCmd.MarkFlagRequired("kvs-arn")
	cloudfrontKeyvaluestoreCmd.AddCommand(cloudfrontKeyvaluestore_getKeyCmd)
}
