package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "This operation tags a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_tagResourceCmd).Standalone()

	dataexchange_tagResourceCmd.Flags().String("resource-arn", "", "An Amazon Resource Name (ARN) that uniquely identifies an AWS resource.")
	dataexchange_tagResourceCmd.Flags().String("tags", "", "A label that consists of a customer-defined key and an optional value.")
	dataexchange_tagResourceCmd.MarkFlagRequired("resource-arn")
	dataexchange_tagResourceCmd.MarkFlagRequired("tags")
	dataexchangeCmd.AddCommand(dataexchange_tagResourceCmd)
}
