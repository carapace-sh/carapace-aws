package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds [tags](https://docs.aws.amazon.com/lambda/latest/dg/tagging.html) to a function, event source mapping, or code signing configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_tagResourceCmd).Standalone()

	lambda_tagResourceCmd.Flags().String("resource", "", "The resource's Amazon Resource Name (ARN).")
	lambda_tagResourceCmd.Flags().String("tags", "", "A list of tags to apply to the resource.")
	lambda_tagResourceCmd.MarkFlagRequired("resource")
	lambda_tagResourceCmd.MarkFlagRequired("tags")
	lambdaCmd.AddCommand(lambda_tagResourceCmd)
}
