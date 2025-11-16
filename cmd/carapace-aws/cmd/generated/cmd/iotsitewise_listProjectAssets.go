package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_listProjectAssetsCmd = &cobra.Command{
	Use:   "list-project-assets",
	Short: "Retrieves a paginated list of assets associated with an IoT SiteWise Monitor project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_listProjectAssetsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_listProjectAssetsCmd).Standalone()

		iotsitewise_listProjectAssetsCmd.Flags().String("max-results", "", "The maximum number of results to return for each paginated request.")
		iotsitewise_listProjectAssetsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
		iotsitewise_listProjectAssetsCmd.Flags().String("project-id", "", "The ID of the project.")
		iotsitewise_listProjectAssetsCmd.MarkFlagRequired("project-id")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_listProjectAssetsCmd)
}
