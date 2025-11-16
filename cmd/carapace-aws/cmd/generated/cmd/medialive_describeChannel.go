package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_describeChannelCmd = &cobra.Command{
	Use:   "describe-channel",
	Short: "Gets details about a channel",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_describeChannelCmd).Standalone()

	medialive_describeChannelCmd.Flags().String("channel-id", "", "channel ID")
	medialive_describeChannelCmd.MarkFlagRequired("channel-id")
	medialiveCmd.AddCommand(medialive_describeChannelCmd)
}
