package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var simspaceweaver_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists all tags on a SimSpace Weaver resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(simspaceweaver_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(simspaceweaver_listTagsForResourceCmd).Standalone()

		simspaceweaver_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
		simspaceweaver_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	simspaceweaverCmd.AddCommand(simspaceweaver_listTagsForResourceCmd)
}
