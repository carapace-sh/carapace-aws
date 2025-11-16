package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var apigateway_getTagsCmd = &cobra.Command{
	Use:   "get-tags",
	Short: "Gets the Tags collection for a given resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(apigateway_getTagsCmd).Standalone()

	apigateway_getTagsCmd.Flags().String("limit", "", "(Not currently supported) The maximum number of returned results per page.")
	apigateway_getTagsCmd.Flags().String("position", "", "(Not currently supported) The current pagination position in the paged result set.")
	apigateway_getTagsCmd.Flags().String("resource-arn", "", "The ARN of a resource that can be tagged.")
	apigateway_getTagsCmd.MarkFlagRequired("resource-arn")
	apigatewayCmd.AddCommand(apigateway_getTagsCmd)
}
