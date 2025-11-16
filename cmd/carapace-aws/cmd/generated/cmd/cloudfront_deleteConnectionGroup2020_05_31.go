package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_deleteConnectionGroup2020_05_31Cmd = &cobra.Command{
	Use:   "delete-connection-group2020_05_31",
	Short: "Deletes a connection group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_deleteConnectionGroup2020_05_31Cmd).Standalone()

	cloudfront_deleteConnectionGroup2020_05_31Cmd.Flags().String("id", "", "The ID of the connection group to delete.")
	cloudfront_deleteConnectionGroup2020_05_31Cmd.Flags().String("if-match", "", "The value of the `ETag` header that you received when retrieving the connection group to delete.")
	cloudfront_deleteConnectionGroup2020_05_31Cmd.MarkFlagRequired("id")
	cloudfront_deleteConnectionGroup2020_05_31Cmd.MarkFlagRequired("if-match")
	cloudfrontCmd.AddCommand(cloudfront_deleteConnectionGroup2020_05_31Cmd)
}
