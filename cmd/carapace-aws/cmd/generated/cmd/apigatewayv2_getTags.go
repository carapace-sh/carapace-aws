package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_getTagsCmd = &cobra.Command{
	Use:   "get-tags",
	Short: "Gets a collection of Tag resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_getTagsCmd).Standalone()

	apigatewayv2_getTagsCmd.Flags().String("resource-arn", "", "The resource ARN for the tag.")
	apigatewayv2_getTagsCmd.MarkFlagRequired("resource-arn")
	apigatewayv2Cmd.AddCommand(apigatewayv2_getTagsCmd)
}
