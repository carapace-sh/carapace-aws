package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Creates a new Tag resource to represent a tag.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_tagResourceCmd).Standalone()

	apigatewayv2_tagResourceCmd.Flags().String("resource-arn", "", "The resource ARN for the tag.")
	apigatewayv2_tagResourceCmd.Flags().String("tags", "", "The collection of tags.")
	apigatewayv2_tagResourceCmd.MarkFlagRequired("resource-arn")
	apigatewayv2Cmd.AddCommand(apigatewayv2_tagResourceCmd)
}
