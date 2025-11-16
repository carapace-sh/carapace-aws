package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehub_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists the tags for your resources in your Resilience Hub applications.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehub_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(resiliencehub_listTagsForResourceCmd).Standalone()

		resiliencehub_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for a specific resource in your Resilience Hub application.")
		resiliencehub_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	resiliencehubCmd.AddCommand(resiliencehub_listTagsForResourceCmd)
}
