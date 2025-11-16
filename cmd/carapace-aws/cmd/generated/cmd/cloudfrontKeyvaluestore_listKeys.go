package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfrontKeyvaluestore_listKeysCmd = &cobra.Command{
	Use:   "list-keys",
	Short: "Returns a list of key value pairs.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfrontKeyvaluestore_listKeysCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfrontKeyvaluestore_listKeysCmd).Standalone()

		cloudfrontKeyvaluestore_listKeysCmd.Flags().String("kvs-arn", "", "The Amazon Resource Name (ARN) of the Key Value Store.")
		cloudfrontKeyvaluestore_listKeysCmd.Flags().String("max-results", "", "Maximum number of results that are returned per call.")
		cloudfrontKeyvaluestore_listKeysCmd.Flags().String("next-token", "", "If nextToken is returned in the response, there are more results available.")
		cloudfrontKeyvaluestore_listKeysCmd.MarkFlagRequired("kvs-arn")
	})
	cloudfrontKeyvaluestoreCmd.AddCommand(cloudfrontKeyvaluestore_listKeysCmd)
}
