package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_untagResourceCmd = &cobra.Command{
	Use:   "untag-resource",
	Short: "Removes a tag from an IoT SiteWise resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_untagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_untagResourceCmd).Standalone()

		iotsitewise_untagResourceCmd.Flags().String("resource-arn", "", "The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the resource to untag.")
		iotsitewise_untagResourceCmd.Flags().String("tag-keys", "", "A list of keys for tags to remove from the resource.")
		iotsitewise_untagResourceCmd.MarkFlagRequired("resource-arn")
		iotsitewise_untagResourceCmd.MarkFlagRequired("tag-keys")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_untagResourceCmd)
}
