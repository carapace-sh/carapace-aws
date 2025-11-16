package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_getCollaborationMlinputChannelCmd = &cobra.Command{
	Use:   "get-collaboration-mlinput-channel",
	Short: "Returns information about a specific ML input channel in a collaboration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_getCollaborationMlinputChannelCmd).Standalone()

	cleanroomsml_getCollaborationMlinputChannelCmd.Flags().String("collaboration-identifier", "", "The collaboration ID of the collaboration that contains the ML input channel that you want to get.")
	cleanroomsml_getCollaborationMlinputChannelCmd.Flags().String("ml-input-channel-arn", "", "The Amazon Resource Name (ARN) of the ML input channel that you want to get.")
	cleanroomsml_getCollaborationMlinputChannelCmd.MarkFlagRequired("collaboration-identifier")
	cleanroomsml_getCollaborationMlinputChannelCmd.MarkFlagRequired("ml-input-channel-arn")
	cleanroomsmlCmd.AddCommand(cleanroomsml_getCollaborationMlinputChannelCmd)
}
