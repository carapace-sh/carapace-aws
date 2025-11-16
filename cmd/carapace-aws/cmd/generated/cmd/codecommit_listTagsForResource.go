package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecommit_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Gets information about Amazon Web Servicestags for a specified Amazon Resource Name (ARN) in CodeCommit.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecommit_listTagsForResourceCmd).Standalone()

	codecommit_listTagsForResourceCmd.Flags().String("next-token", "", "An enumeration token that, when provided in a request, returns the next batch of the results.")
	codecommit_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource for which you want to get information about tags, if any.")
	codecommit_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	codecommitCmd.AddCommand(codecommit_listTagsForResourceCmd)
}
