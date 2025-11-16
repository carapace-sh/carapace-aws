package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudfront_updateKeyGroup2020_05_31Cmd = &cobra.Command{
	Use:   "update-key-group2020_05_31",
	Short: "Updates a key group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudfront_updateKeyGroup2020_05_31Cmd).Standalone()

	cloudfront_updateKeyGroup2020_05_31Cmd.Flags().String("id", "", "The identifier of the key group that you are updating.")
	cloudfront_updateKeyGroup2020_05_31Cmd.Flags().String("if-match", "", "The version of the key group that you are updating.")
	cloudfront_updateKeyGroup2020_05_31Cmd.Flags().String("key-group-config", "", "The key group configuration.")
	cloudfront_updateKeyGroup2020_05_31Cmd.MarkFlagRequired("id")
	cloudfront_updateKeyGroup2020_05_31Cmd.MarkFlagRequired("key-group-config")
	cloudfrontCmd.AddCommand(cloudfront_updateKeyGroup2020_05_31Cmd)
}
