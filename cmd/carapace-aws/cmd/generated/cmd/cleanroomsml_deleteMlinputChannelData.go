package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_deleteMlinputChannelDataCmd = &cobra.Command{
	Use:   "delete-mlinput-channel-data",
	Short: "Provides the information necessary to delete an ML input channel.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_deleteMlinputChannelDataCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_deleteMlinputChannelDataCmd).Standalone()

		cleanroomsml_deleteMlinputChannelDataCmd.Flags().String("membership-identifier", "", "The membership ID of the membership that contains the ML input channel you want to delete.")
		cleanroomsml_deleteMlinputChannelDataCmd.Flags().String("ml-input-channel-arn", "", "The Amazon Resource Name (ARN) of the ML input channel that you want to delete.")
		cleanroomsml_deleteMlinputChannelDataCmd.MarkFlagRequired("membership-identifier")
		cleanroomsml_deleteMlinputChannelDataCmd.MarkFlagRequired("ml-input-channel-arn")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_deleteMlinputChannelDataCmd)
}
