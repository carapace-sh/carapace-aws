package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ram_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds the specified tag keys and values to a resource share or managed permission.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ram_tagResourceCmd).Standalone()

	ram_tagResourceCmd.Flags().String("resource-arn", "", "Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the managed permission that you want to add tags to.")
	ram_tagResourceCmd.Flags().String("resource-share-arn", "", "Specifies the [Amazon Resource Name (ARN)](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the resource share that you want to add tags to.")
	ram_tagResourceCmd.Flags().String("tags", "", "A list of one or more tag key and value pairs.")
	ram_tagResourceCmd.MarkFlagRequired("tags")
	ramCmd.AddCommand(ram_tagResourceCmd)
}
