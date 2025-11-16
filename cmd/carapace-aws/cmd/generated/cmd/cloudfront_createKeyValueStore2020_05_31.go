package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_createKeyValueStore2020_05_31Cmd = &cobra.Command{
	Use:   "create-key-value-store2020_05_31",
	Short: "Specifies the key value store resource to add to your account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_createKeyValueStore2020_05_31Cmd).Standalone()

	cloudfront_createKeyValueStore2020_05_31Cmd.Flags().String("comment", "", "The comment of the key value store.")
	cloudfront_createKeyValueStore2020_05_31Cmd.Flags().String("import-source", "", "The S3 bucket that provides the source for the import.")
	cloudfront_createKeyValueStore2020_05_31Cmd.Flags().String("name", "", "The name of the key value store.")
	cloudfront_createKeyValueStore2020_05_31Cmd.MarkFlagRequired("name")
	cloudfrontCmd.AddCommand(cloudfront_createKeyValueStore2020_05_31Cmd)
}
