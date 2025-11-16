package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassv2_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag from an IoT Greengrass resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassv2_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrassv2_untagResourceCmd).Standalone()

		greengrassv2_untagResourceCmd.Flags().String("resource-arn", "", "The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the resource to untag.")
		greengrassv2_untagResourceCmd.Flags().String("tag-keys", "", "A list of keys for tags to remove from the resource.")
		greengrassv2_untagResourceCmd.MarkFlagRequired("resource-arn")
		greengrassv2_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	greengrassv2Cmd.AddCommand(greengrassv2_untagResourceCmd)
}
