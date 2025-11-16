package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_listPortalsCmd = &cobra.Command{
	Use:   "list-portals",
	Short: "Retrieves a paginated list of IoT SiteWise Monitor portals.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_listPortalsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_listPortalsCmd).Standalone()

		iotsitewise_listPortalsCmd.Flags().String("max-results", "", "The maximum number of results to return for each paginated request.")
		iotsitewise_listPortalsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_listPortalsCmd)
}
