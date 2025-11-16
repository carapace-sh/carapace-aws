package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "This operation lists the tags on the resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(dataexchange_listTagsForResourceCmd).Standalone()

		dataexchange_listTagsForResourceCmd.Flags().String("resource-arn", "", "An Amazon Resource Name (ARN) that uniquely identifies an AWS resource.")
		dataexchange_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	dataexchangeCmd.AddCommand(dataexchange_listTagsForResourceCmd)
}
