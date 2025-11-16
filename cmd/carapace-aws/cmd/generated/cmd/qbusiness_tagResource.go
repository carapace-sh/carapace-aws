package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var qbusiness_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds the specified tag to the specified Amazon Q Business application or data source resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(qbusiness_tagResourceCmd).Standalone()

	qbusiness_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the Amazon Q Business application or data source to tag.")
	qbusiness_tagResourceCmd.Flags().String("tags", "", "A list of tag keys to add to the Amazon Q Business application or data source.")
	qbusiness_tagResourceCmd.MarkFlagRequired("resource-arn")
	qbusiness_tagResourceCmd.MarkFlagRequired("tags")
	qbusinessCmd.AddCommand(qbusiness_tagResourceCmd)
}
