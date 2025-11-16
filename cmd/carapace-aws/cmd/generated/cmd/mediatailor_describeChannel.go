package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediatailor_describeChannelCmd = &cobra.Command{
	Use:   "describe-channel",
	Short: "Describes a channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediatailor_describeChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediatailor_describeChannelCmd).Standalone()

		mediatailor_describeChannelCmd.Flags().String("channel-name", "", "The name of the channel.")
		mediatailor_describeChannelCmd.MarkFlagRequired("channel-name")
	})
	mediatailorCmd.AddCommand(mediatailor_describeChannelCmd)
}
