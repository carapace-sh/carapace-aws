package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigatewayv2_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Deletes a Tag.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigatewayv2_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigatewayv2_untagResourceCmd).Standalone()

		apigatewayv2_untagResourceCmd.Flags().String("resource-arn", "", "The resource ARN for the tag.")
		apigatewayv2_untagResourceCmd.Flags().String("tag-keys", "", "The Tag keys to delete")
		apigatewayv2_untagResourceCmd.MarkFlagRequired("resource-arn")
		apigatewayv2_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	apigatewayv2Cmd.AddCommand(apigatewayv2_untagResourceCmd)
}
