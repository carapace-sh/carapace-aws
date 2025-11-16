package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_updateKeyValueStore2020_05_31Cmd = &cobra.Command{
	Use:   "update-key-value-store2020_05_31",
	Short: "Specifies the key value store to update.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_updateKeyValueStore2020_05_31Cmd).Standalone()

	cloudfront_updateKeyValueStore2020_05_31Cmd.Flags().String("comment", "", "The comment of the key value store to update.")
	cloudfront_updateKeyValueStore2020_05_31Cmd.Flags().String("if-match", "", "The key value store to update, if a match occurs.")
	cloudfront_updateKeyValueStore2020_05_31Cmd.Flags().String("name", "", "The name of the key value store to update.")
	cloudfront_updateKeyValueStore2020_05_31Cmd.MarkFlagRequired("comment")
	cloudfront_updateKeyValueStore2020_05_31Cmd.MarkFlagRequired("if-match")
	cloudfront_updateKeyValueStore2020_05_31Cmd.MarkFlagRequired("name")
	cloudfrontCmd.AddCommand(cloudfront_updateKeyValueStore2020_05_31Cmd)
}
