package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_getConnectionGroup2020_05_31Cmd = &cobra.Command{
	Use:   "get-connection-group2020_05_31",
	Short: "Gets information about a connection group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_getConnectionGroup2020_05_31Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudfront_getConnectionGroup2020_05_31Cmd).Standalone()

		cloudfront_getConnectionGroup2020_05_31Cmd.Flags().String("identifier", "", "The ID, name, or Amazon Resource Name (ARN) of the connection group.")
		cloudfront_getConnectionGroup2020_05_31Cmd.MarkFlagRequired("identifier")
	})
	cloudfrontCmd.AddCommand(cloudfront_getConnectionGroup2020_05_31Cmd)
}
