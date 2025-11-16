package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53profiles_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds one or more tags to a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53profiles_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53profiles_tagResourceCmd).Standalone()

		route53profiles_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for the resource that you want to add tags to.")
		route53profiles_tagResourceCmd.Flags().String("tags", "", "The tags that you want to add to the specified resource.")
		route53profiles_tagResourceCmd.MarkFlagRequired("resource-arn")
		route53profiles_tagResourceCmd.MarkFlagRequired("tags")
	})
	route53profilesCmd.AddCommand(route53profiles_tagResourceCmd)
}
