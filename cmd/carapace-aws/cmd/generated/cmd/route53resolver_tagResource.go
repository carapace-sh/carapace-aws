package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53resolver_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds one or more tags to a specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53resolver_tagResourceCmd).Standalone()

	route53resolver_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) for the resource that you want to add tags to.")
	route53resolver_tagResourceCmd.Flags().String("tags", "", "The tags that you want to add to the specified resource.")
	route53resolver_tagResourceCmd.MarkFlagRequired("resource-arn")
	route53resolver_tagResourceCmd.MarkFlagRequired("tags")
	route53resolverCmd.AddCommand(route53resolver_tagResourceCmd)
}
