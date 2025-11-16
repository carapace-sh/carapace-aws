package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds or updates a tag on a given resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_tagResourceCmd).Standalone()

	apigateway_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of a resource that can be tagged.")
	apigateway_tagResourceCmd.Flags().String("tags", "", "The key-value map of strings.")
	apigateway_tagResourceCmd.MarkFlagRequired("resource-arn")
	apigateway_tagResourceCmd.MarkFlagRequired("tags")
	apigatewayCmd.AddCommand(apigateway_tagResourceCmd)
}
