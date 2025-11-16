package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_getMlinputChannelCmd = &cobra.Command{
	Use:   "get-mlinput-channel",
	Short: "Returns information about an ML input channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_getMlinputChannelCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_getMlinputChannelCmd).Standalone()

		cleanroomsml_getMlinputChannelCmd.Flags().String("membership-identifier", "", "The membership ID of the membership that contains the ML input channel that you want to get.")
		cleanroomsml_getMlinputChannelCmd.Flags().String("ml-input-channel-arn", "", "The Amazon Resource Name (ARN) of the ML input channel that you want to get.")
		cleanroomsml_getMlinputChannelCmd.MarkFlagRequired("membership-identifier")
		cleanroomsml_getMlinputChannelCmd.MarkFlagRequired("ml-input-channel-arn")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_getMlinputChannelCmd)
}
