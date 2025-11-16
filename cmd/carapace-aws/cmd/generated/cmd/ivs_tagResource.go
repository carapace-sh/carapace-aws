package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ivs_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds or updates tags for the Amazon Web Services resource with the specified ARN.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ivs_tagResourceCmd).Standalone()

	ivs_tagResourceCmd.Flags().String("resource-arn", "", "ARN of the resource for which tags are to be added or updated.")
	ivs_tagResourceCmd.Flags().String("tags", "", "Array of tags to be added or updated.")
	ivs_tagResourceCmd.MarkFlagRequired("resource-arn")
	ivs_tagResourceCmd.MarkFlagRequired("tags")
	ivsCmd.AddCommand(ivs_tagResourceCmd)
}
