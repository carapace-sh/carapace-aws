package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_getKeyGroup2020_05_31Cmd = &cobra.Command{
	Use:   "get-key-group2020_05_31",
	Short: "Gets a key group, including the date and time when the key group was last modified.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_getKeyGroup2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_getKeyGroup2020_05_31Cmd).Standalone()

		cloudfront_getKeyGroup2020_05_31Cmd.Flags().String("id", "", "The identifier of the key group that you are getting.")
		cloudfront_getKeyGroup2020_05_31Cmd.MarkFlagRequired("id")
	})
	cloudfrontCmd.AddCommand(cloudfront_getKeyGroup2020_05_31Cmd)
}
