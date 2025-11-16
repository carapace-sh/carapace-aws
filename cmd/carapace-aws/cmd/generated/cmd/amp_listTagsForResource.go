package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var amp_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "The `ListTagsForResource` operation returns the tags that are associated with an Amazon Managed Service for Prometheus resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(amp_listTagsForResourceCmd).Standalone()

	amp_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource to list tages for.")
	amp_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	ampCmd.AddCommand(amp_listTagsForResourceCmd)
}
