package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_describeKeyValueStore2020_05_31Cmd = &cobra.Command{
	Use:   "describe-key-value-store2020_05_31",
	Short: "Specifies the key value store and its configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_describeKeyValueStore2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_describeKeyValueStore2020_05_31Cmd).Standalone()

		cloudfront_describeKeyValueStore2020_05_31Cmd.Flags().String("name", "", "The name of the key value store.")
		cloudfront_describeKeyValueStore2020_05_31Cmd.MarkFlagRequired("name")
	})
	cloudfrontCmd.AddCommand(cloudfront_describeKeyValueStore2020_05_31Cmd)
}
