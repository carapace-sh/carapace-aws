package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfrontKeyvaluestore_deleteKeyCmd = &cobra.Command{
	Use:   "delete-key",
	Short: "Deletes the key value pair specified by the key.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfrontKeyvaluestore_deleteKeyCmd).Standalone()

	cloudfrontKeyvaluestore_deleteKeyCmd.Flags().String("if-match", "", "The current version (ETag) of the Key Value Store that you are deleting keys from, which you can get using DescribeKeyValueStore.")
	cloudfrontKeyvaluestore_deleteKeyCmd.Flags().String("key", "", "The key to delete.")
	cloudfrontKeyvaluestore_deleteKeyCmd.Flags().String("kvs-arn", "", "The Amazon Resource Name (ARN) of the Key Value Store.")
	cloudfrontKeyvaluestore_deleteKeyCmd.MarkFlagRequired("if-match")
	cloudfrontKeyvaluestore_deleteKeyCmd.MarkFlagRequired("key")
	cloudfrontKeyvaluestore_deleteKeyCmd.MarkFlagRequired("kvs-arn")
	cloudfrontKeyvaluestoreCmd.AddCommand(cloudfrontKeyvaluestore_deleteKeyCmd)
}
