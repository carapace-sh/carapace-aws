package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_deleteKeyGroup2020_05_31Cmd = &cobra.Command{
	Use:   "delete-key-group2020_05_31",
	Short: "Deletes a key group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_deleteKeyGroup2020_05_31Cmd).Standalone()

	cloudfront_deleteKeyGroup2020_05_31Cmd.Flags().String("id", "", "The identifier of the key group that you are deleting.")
	cloudfront_deleteKeyGroup2020_05_31Cmd.Flags().String("if-match", "", "The version of the key group that you are deleting.")
	cloudfront_deleteKeyGroup2020_05_31Cmd.MarkFlagRequired("id")
	cloudfrontCmd.AddCommand(cloudfront_deleteKeyGroup2020_05_31Cmd)
}
