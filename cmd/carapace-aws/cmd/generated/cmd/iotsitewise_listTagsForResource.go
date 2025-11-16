package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotsitewise_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Retrieves the list of tags for an IoT SiteWise resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotsitewise_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotsitewise_listTagsForResourceCmd).Standalone()

		iotsitewise_listTagsForResourceCmd.Flags().String("resource-arn", "", "The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the resource.")
		iotsitewise_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	iotsitewiseCmd.AddCommand(iotsitewise_listTagsForResourceCmd)
}
