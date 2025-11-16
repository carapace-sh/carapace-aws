package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag from a given resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(apigateway_untagResourceCmd).Standalone()

		apigateway_untagResourceCmd.Flags().String("resource-arn", "", "The ARN of a resource that can be tagged.")
		apigateway_untagResourceCmd.Flags().String("tag-keys", "", "The Tag keys to delete.")
		apigateway_untagResourceCmd.MarkFlagRequired("resource-arn")
		apigateway_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	apigatewayCmd.AddCommand(apigateway_untagResourceCmd)
}
