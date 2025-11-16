package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ce_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "An API operation for adding one or more tags (key-value pairs) to a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ce_tagResourceCmd).Standalone()

	ce_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource.")
	ce_tagResourceCmd.Flags().String("resource-tags", "", "A list of tag key-value pairs to be added to the resource.")
	ce_tagResourceCmd.MarkFlagRequired("resource-arn")
	ce_tagResourceCmd.MarkFlagRequired("resource-tags")
	ceCmd.AddCommand(ce_tagResourceCmd)
}
