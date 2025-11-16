package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourcegroupstaggingapi_tagResourcesCmd = &cobra.Command{
	Use:   "tag-resources",
	Short: "Applies one or more tags to the specified resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourcegroupstaggingapi_tagResourcesCmd).Standalone()

	resourcegroupstaggingapi_tagResourcesCmd.Flags().String("resource-arnlist", "", "Specifies the list of ARNs of the resources that you want to apply tags to.")
	resourcegroupstaggingapi_tagResourcesCmd.Flags().String("tags", "", "Specifies a list of tags that you want to add to the specified resources.")
	resourcegroupstaggingapi_tagResourcesCmd.MarkFlagRequired("resource-arnlist")
	resourcegroupstaggingapi_tagResourcesCmd.MarkFlagRequired("tags")
	resourcegroupstaggingapiCmd.AddCommand(resourcegroupstaggingapi_tagResourcesCmd)
}
