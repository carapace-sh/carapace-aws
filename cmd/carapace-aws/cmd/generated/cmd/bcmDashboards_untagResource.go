package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bcmDashboards_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes specified tags from a dashboard resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bcmDashboards_untagResourceCmd).Standalone()

	bcmDashboards_untagResourceCmd.Flags().String("resource-arn", "", "The unique identifier for the resource.")
	bcmDashboards_untagResourceCmd.Flags().String("resource-tag-keys", "", "The keys of the tags to remove from the dashboard resource.")
	bcmDashboards_untagResourceCmd.MarkFlagRequired("resource-arn")
	bcmDashboards_untagResourceCmd.MarkFlagRequired("resource-tag-keys")
	bcmDashboardsCmd.AddCommand(bcmDashboards_untagResourceCmd)
}
