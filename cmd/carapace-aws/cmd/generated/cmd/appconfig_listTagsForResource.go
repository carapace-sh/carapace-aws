package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfig_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Retrieves the list of key-value tags assigned to the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfig_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appconfig_listTagsForResourceCmd).Standalone()

		appconfig_listTagsForResourceCmd.Flags().String("resource-arn", "", "The resource ARN.")
		appconfig_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	appconfigCmd.AddCommand(appconfig_listTagsForResourceCmd)
}
