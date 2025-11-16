package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dataexchange_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "This operation removes one or more tags from a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dataexchange_untagResourceCmd).Standalone()

	dataexchange_untagResourceCmd.Flags().String("resource-arn", "", "An Amazon Resource Name (ARN) that uniquely identifies an AWS resource.")
	dataexchange_untagResourceCmd.Flags().String("tag-keys", "", "The key tags.")
	dataexchange_untagResourceCmd.MarkFlagRequired("resource-arn")
	dataexchange_untagResourceCmd.MarkFlagRequired("tag-keys")
	dataexchangeCmd.AddCommand(dataexchange_untagResourceCmd)
}
