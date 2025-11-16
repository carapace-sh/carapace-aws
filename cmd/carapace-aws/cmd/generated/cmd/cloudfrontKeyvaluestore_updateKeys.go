package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfrontKeyvaluestore_updateKeysCmd = &cobra.Command{
	Use:   "update-keys",
	Short: "Puts or Deletes multiple key value pairs in a single, all-or-nothing operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfrontKeyvaluestore_updateKeysCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfrontKeyvaluestore_updateKeysCmd).Standalone()

		cloudfrontKeyvaluestore_updateKeysCmd.Flags().String("deletes", "", "List of keys to delete.")
		cloudfrontKeyvaluestore_updateKeysCmd.Flags().String("if-match", "", "The current version (ETag) of the Key Value Store that you are updating keys of, which you can get using DescribeKeyValueStore.")
		cloudfrontKeyvaluestore_updateKeysCmd.Flags().String("kvs-arn", "", "The Amazon Resource Name (ARN) of the Key Value Store.")
		cloudfrontKeyvaluestore_updateKeysCmd.Flags().String("puts", "", "List of key value pairs to put.")
		cloudfrontKeyvaluestore_updateKeysCmd.MarkFlagRequired("if-match")
		cloudfrontKeyvaluestore_updateKeysCmd.MarkFlagRequired("kvs-arn")
	})
	cloudfrontKeyvaluestoreCmd.AddCommand(cloudfrontKeyvaluestore_updateKeysCmd)
}
