package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cleanroomsml_listMlinputChannelsCmd = &cobra.Command{
	Use:   "list-mlinput-channels",
	Short: "Returns a list of ML input channels.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cleanroomsml_listMlinputChannelsCmd).Standalone()

	cleanroomsml_listMlinputChannelsCmd.Flags().String("max-results", "", "The maximum number of ML input channels to return.")
	cleanroomsml_listMlinputChannelsCmd.Flags().String("membership-identifier", "", "The membership ID of the membership that contains the ML input channels that you want to list.")
	cleanroomsml_listMlinputChannelsCmd.Flags().String("next-token", "", "The token value retrieved from a previous call to access the next page of results.")
	cleanroomsml_listMlinputChannelsCmd.MarkFlagRequired("membership-identifier")
	cleanroomsmlCmd.AddCommand(cleanroomsml_listMlinputChannelsCmd)
}
