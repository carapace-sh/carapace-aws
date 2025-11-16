package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_stopChannelCmd = &cobra.Command{
	Use:   "stop-channel",
	Short: "Stops a running channel",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_stopChannelCmd).Standalone()

	medialive_stopChannelCmd.Flags().String("channel-id", "", "A request to stop a running channel")
	medialive_stopChannelCmd.MarkFlagRequired("channel-id")
	medialiveCmd.AddCommand(medialive_stopChannelCmd)
}
