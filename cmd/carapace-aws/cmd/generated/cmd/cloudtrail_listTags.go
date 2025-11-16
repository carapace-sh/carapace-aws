package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_listTagsCmd = &cobra.Command{
	Use:   "list-tags",
	Short: "Lists the tags for the specified trails, event data stores, dashboards, or channels in the current Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_listTagsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudtrail_listTagsCmd).Standalone()

		cloudtrail_listTagsCmd.Flags().String("next-token", "", "Reserved for future use.")
		cloudtrail_listTagsCmd.Flags().String("resource-id-list", "", "Specifies a list of trail, event data store, dashboard, or channel ARNs whose tags will be listed.")
		cloudtrail_listTagsCmd.MarkFlagRequired("resource-id-list")
	})
	cloudtrailCmd.AddCommand(cloudtrail_listTagsCmd)
}
