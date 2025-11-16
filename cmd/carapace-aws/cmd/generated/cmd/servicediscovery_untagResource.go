package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscovery_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscovery_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(servicediscovery_untagResourceCmd).Standalone()

		servicediscovery_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to retrieve tags for.")
		servicediscovery_untagResourceCmd.Flags().String("tag-keys", "", "The tag keys to remove from the specified resource.")
		servicediscovery_untagResourceCmd.MarkFlagRequired("resource-arn")
		servicediscovery_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	servicediscoveryCmd.AddCommand(servicediscovery_untagResourceCmd)
}
