package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds tags to an IoT SiteWise resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_tagResourceCmd).Standalone()

		iotsitewise_tagResourceCmd.Flags().String("resource-arn", "", "The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the resource to tag.")
		iotsitewise_tagResourceCmd.Flags().String("tags", "", "A list of key-value pairs that contain metadata for the resource.")
		iotsitewise_tagResourceCmd.MarkFlagRequired("resource-arn")
		iotsitewise_tagResourceCmd.MarkFlagRequired("tags")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_tagResourceCmd)
}
