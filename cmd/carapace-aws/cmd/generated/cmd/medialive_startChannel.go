package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_startChannelCmd = &cobra.Command{
	Use:   "start-channel",
	Short: "Starts an existing channel",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_startChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_startChannelCmd).Standalone()

		medialive_startChannelCmd.Flags().String("channel-id", "", "A request to start a channel")
		medialive_startChannelCmd.MarkFlagRequired("channel-id")
	})
	medialiveCmd.AddCommand(medialive_startChannelCmd)
}
