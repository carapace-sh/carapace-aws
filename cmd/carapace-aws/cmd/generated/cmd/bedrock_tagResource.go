package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var bedrock_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Associate tags with a resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(bedrock_tagResourceCmd).Standalone()

	bedrock_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the resource to tag.")
	bedrock_tagResourceCmd.Flags().String("tags", "", "Tags to associate with the resource.")
	bedrock_tagResourceCmd.MarkFlagRequired("resource-arn")
	bedrock_tagResourceCmd.MarkFlagRequired("tags")
	bedrockCmd.AddCommand(bedrock_tagResourceCmd)
}
