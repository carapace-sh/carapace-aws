package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lambda_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes [tags](https://docs.aws.amazon.com/lambda/latest/dg/tagging.html) from a function, event source mapping, or code signing configuration.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lambda_untagResourceCmd).Standalone()

	lambda_untagResourceCmd.Flags().String("resource", "", "The resource's Amazon Resource Name (ARN).")
	lambda_untagResourceCmd.Flags().String("tag-keys", "", "A list of tag keys to remove from the resource.")
	lambda_untagResourceCmd.MarkFlagRequired("resource")
	lambda_untagResourceCmd.MarkFlagRequired("tag-keys")
	lambdaCmd.AddCommand(lambda_untagResourceCmd)
}
