package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscovery_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Lists tags for the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscovery_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicediscovery_listTagsForResourceCmd).Standalone()

		servicediscovery_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to retrieve tags for.")
		servicediscovery_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	servicediscoveryCmd.AddCommand(servicediscovery_listTagsForResourceCmd)
}
