package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var servicediscovery_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds one or more tags to the specified resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(servicediscovery_tagResourceCmd).Standalone()

	servicediscovery_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource that you want to retrieve tags for.")
	servicediscovery_tagResourceCmd.Flags().String("tags", "", "The tags to add to the specified resource.")
	servicediscovery_tagResourceCmd.MarkFlagRequired("resource-arn")
	servicediscovery_tagResourceCmd.MarkFlagRequired("tags")
	servicediscoveryCmd.AddCommand(servicediscovery_tagResourceCmd)
}
