package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediapackage_describeChannelCmd = &cobra.Command{
	Use:   "describe-channel",
	Short: "Gets details about a Channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediapackage_describeChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediapackage_describeChannelCmd).Standalone()

		mediapackage_describeChannelCmd.Flags().String("id", "", "The ID of a Channel.")
		mediapackage_describeChannelCmd.MarkFlagRequired("id")
	})
	mediapackageCmd.AddCommand(mediapackage_describeChannelCmd)
}
