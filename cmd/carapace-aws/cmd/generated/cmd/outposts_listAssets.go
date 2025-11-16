package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var outposts_listAssetsCmd = &cobra.Command{
	Use:   "list-assets",
	Short: "Lists the hardware assets for the specified Outpost.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(outposts_listAssetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(outposts_listAssetsCmd).Standalone()

		outposts_listAssetsCmd.Flags().String("host-id-filter", "", "Filters the results by the host ID of a Dedicated Host.")
		outposts_listAssetsCmd.Flags().String("max-results", "", "")
		outposts_listAssetsCmd.Flags().String("next-token", "", "")
		outposts_listAssetsCmd.Flags().String("outpost-identifier", "", "The ID or the Amazon Resource Name (ARN) of the Outpost.")
		outposts_listAssetsCmd.Flags().String("status-filter", "", "Filters the results by state.")
		outposts_listAssetsCmd.MarkFlagRequired("outpost-identifier")
	})
	outpostsCmd.AddCommand(outposts_listAssetsCmd)
}
