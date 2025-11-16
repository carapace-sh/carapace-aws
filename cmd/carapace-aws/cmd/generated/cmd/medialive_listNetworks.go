package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_listNetworksCmd = &cobra.Command{
	Use:   "list-networks",
	Short: "Retrieve the list of Networks.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_listNetworksCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_listNetworksCmd).Standalone()

		medialive_listNetworksCmd.Flags().String("max-results", "", "The maximum number of items to return.")
		medialive_listNetworksCmd.Flags().String("next-token", "", "The token to retrieve the next page of results.")
	})
	medialiveCmd.AddCommand(medialive_listNetworksCmd)
}
