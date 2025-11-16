package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmDashboards_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of all tags associated with a specified dashboard resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmDashboards_listTagsForResourceCmd).Standalone()

	bcmDashboards_listTagsForResourceCmd.Flags().String("resource-arn", "", "The unique identifier for the resource.")
	bcmDashboards_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	bcmDashboardsCmd.AddCommand(bcmDashboards_listTagsForResourceCmd)
}
