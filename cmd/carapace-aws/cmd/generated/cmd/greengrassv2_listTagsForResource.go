package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var greengrassv2_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Retrieves the list of tags for an IoT Greengrass resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(greengrassv2_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(greengrassv2_listTagsForResourceCmd).Standalone()

		greengrassv2_listTagsForResourceCmd.Flags().String("resource-arn", "", "The [ARN](https://docs.aws.amazon.com/general/latest/gr/aws-arns-and-namespaces.html) of the resource.")
		greengrassv2_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	greengrassv2Cmd.AddCommand(greengrassv2_listTagsForResourceCmd)
}
