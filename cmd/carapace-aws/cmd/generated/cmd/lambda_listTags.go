package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_listTagsCmd = &cobra.Command{
	Use:   "list-tags",
	Short: "Returns a function, event source mapping, or code signing configuration's [tags](https://docs.aws.amazon.com/lambda/latest/dg/tagging.html). You can also view function tags with [GetFunction]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_listTagsCmd).Standalone()

	lambda_listTagsCmd.Flags().String("resource", "", "The resource's Amazon Resource Name (ARN).")
	lambda_listTagsCmd.MarkFlagRequired("resource")
	lambdaCmd.AddCommand(lambda_listTagsCmd)
}
