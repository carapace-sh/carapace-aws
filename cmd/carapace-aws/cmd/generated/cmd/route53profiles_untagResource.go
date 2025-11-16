package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53profiles_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes one or more tags from a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53profiles_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53profiles_untagResourceCmd).Standalone()

		route53profiles_untagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for the resource that you want to remove tags from.")
		route53profiles_untagResourceCmd.Flags().String("tag-keys", "", "The tags that you want to remove to the specified resource.")
		route53profiles_untagResourceCmd.MarkFlagRequired("resource-arn")
		route53profiles_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	route53profilesCmd.AddCommand(route53profiles_untagResourceCmd)
}
