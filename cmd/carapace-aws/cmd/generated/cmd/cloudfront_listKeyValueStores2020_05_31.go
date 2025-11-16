package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_listKeyValueStores2020_05_31Cmd = &cobra.Command{
	Use:   "list-key-value-stores2020_05_31",
	Short: "Specifies the key value stores to list.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_listKeyValueStores2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_listKeyValueStores2020_05_31Cmd).Standalone()

		cloudfront_listKeyValueStores2020_05_31Cmd.Flags().String("marker", "", "The marker associated with the key value stores list.")
		cloudfront_listKeyValueStores2020_05_31Cmd.Flags().String("max-items", "", "The maximum number of items in the key value stores list.")
		cloudfront_listKeyValueStores2020_05_31Cmd.Flags().String("status", "", "The status of the request for the key value stores list.")
	})
	cloudfrontCmd.AddCommand(cloudfront_listKeyValueStores2020_05_31Cmd)
}
