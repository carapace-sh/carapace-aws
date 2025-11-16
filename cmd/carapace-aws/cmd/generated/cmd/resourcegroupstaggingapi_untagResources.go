package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resourcegroupstaggingapi_untagResourcesCmd = &cobra.Command{
	Use:   "untag-resources",
	Short: "Removes the specified tags from the specified resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resourcegroupstaggingapi_untagResourcesCmd).Standalone()

	resourcegroupstaggingapi_untagResourcesCmd.Flags().String("resource-arnlist", "", "Specifies a list of ARNs of the resources that you want to remove tags from.")
	resourcegroupstaggingapi_untagResourcesCmd.Flags().String("tag-keys", "", "Specifies a list of tag keys that you want to remove from the specified resources.")
	resourcegroupstaggingapi_untagResourcesCmd.MarkFlagRequired("resource-arnlist")
	resourcegroupstaggingapi_untagResourcesCmd.MarkFlagRequired("tag-keys")
	resourcegroupstaggingapiCmd.AddCommand(resourcegroupstaggingapi_untagResourcesCmd)
}
