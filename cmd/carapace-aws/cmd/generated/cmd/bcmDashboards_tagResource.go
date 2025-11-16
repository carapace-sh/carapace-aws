package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmDashboards_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds or updates tags for a specified dashboard resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmDashboards_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(bcmDashboards_tagResourceCmd).Standalone()

		bcmDashboards_tagResourceCmd.Flags().String("resource-arn", "", "The unique identifier for the resource.")
		bcmDashboards_tagResourceCmd.Flags().String("resource-tags", "", "The tags to add to the dashboard resource.")
		bcmDashboards_tagResourceCmd.MarkFlagRequired("resource-arn")
		bcmDashboards_tagResourceCmd.MarkFlagRequired("resource-tags")
	})
	bcmDashboardsCmd.AddCommand(bcmDashboards_tagResourceCmd)
}
