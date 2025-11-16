package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appflow_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Retrieves the tags that are associated with a specified flow.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appflow_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(appflow_listTagsForResourceCmd).Standalone()

		appflow_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the specified flow.")
		appflow_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	appflowCmd.AddCommand(appflow_listTagsForResourceCmd)
}
