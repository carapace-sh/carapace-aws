package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_listProjectsCmd = &cobra.Command{
	Use:   "list-projects",
	Short: "Retrieves a paginated list of projects for an IoT SiteWise Monitor portal.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_listProjectsCmd).Standalone()

	iotsitewise_listProjectsCmd.Flags().String("max-results", "", "The maximum number of results to return for each paginated request.")
	iotsitewise_listProjectsCmd.Flags().String("next-token", "", "The token to be used for the next set of paginated results.")
	iotsitewise_listProjectsCmd.Flags().String("portal-id", "", "The ID of the portal.")
	iotsitewise_listProjectsCmd.MarkFlagRequired("portal-id")
	iotsitewiseCmd.AddCommand(iotsitewise_listProjectsCmd)
}
