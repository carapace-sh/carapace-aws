package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassv2_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds tags to an IoT Greengrass resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassv2_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrassv2_tagResourceCmd).Standalone()

		greengrassv2_tagResourceCmd.Flags().String("resource-arn", "", "The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the resource to tag.")
		greengrassv2_tagResourceCmd.Flags().String("tags", "", "A list of key-value pairs that contain metadata for the resource.")
		greengrassv2_tagResourceCmd.MarkFlagRequired("resource-arn")
		greengrassv2_tagResourceCmd.MarkFlagRequired("tags")
	})
	greengrassv2Cmd.AddCommand(greengrassv2_tagResourceCmd)
}
