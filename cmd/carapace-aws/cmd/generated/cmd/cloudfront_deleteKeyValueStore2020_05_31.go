package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_deleteKeyValueStore2020_05_31Cmd = &cobra.Command{
	Use:   "delete-key-value-store2020_05_31",
	Short: "Specifies the key value store to delete.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_deleteKeyValueStore2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_deleteKeyValueStore2020_05_31Cmd).Standalone()

		cloudfront_deleteKeyValueStore2020_05_31Cmd.Flags().String("if-match", "", "The key value store to delete, if a match occurs.")
		cloudfront_deleteKeyValueStore2020_05_31Cmd.Flags().String("name", "", "The name of the key value store.")
		cloudfront_deleteKeyValueStore2020_05_31Cmd.MarkFlagRequired("if-match")
		cloudfront_deleteKeyValueStore2020_05_31Cmd.MarkFlagRequired("name")
	})
	cloudfrontCmd.AddCommand(cloudfront_deleteKeyValueStore2020_05_31Cmd)
}
