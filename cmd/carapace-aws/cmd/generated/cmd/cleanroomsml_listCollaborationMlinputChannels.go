package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_listCollaborationMlinputChannelsCmd = &cobra.Command{
	Use:   "list-collaboration-mlinput-channels",
	Short: "Returns a list of the ML input channels in a collaboration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_listCollaborationMlinputChannelsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cleanroomsml_listCollaborationMlinputChannelsCmd).Standalone()

		cleanroomsml_listCollaborationMlinputChannelsCmd.Flags().String("collaboration-identifier", "", "The collaboration ID of the collaboration that contains the ML input channels that you want to list.")
		cleanroomsml_listCollaborationMlinputChannelsCmd.Flags().String("max-results", "", "The maximum number of results to return.")
		cleanroomsml_listCollaborationMlinputChannelsCmd.Flags().String("next-token", "", "The token value retrieved from a previous call to access the next page of results.")
		cleanroomsml_listCollaborationMlinputChannelsCmd.MarkFlagRequired("collaboration-identifier")
	})
	cleanroomsmlCmd.AddCommand(cleanroomsml_listCollaborationMlinputChannelsCmd)
}
