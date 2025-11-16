package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "List all tags for your Elastic Disaster Recovery resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_listTagsForResourceCmd).Standalone()

	drs_listTagsForResourceCmd.Flags().String("resource-arn", "", "The ARN of the resource whose tags should be returned.")
	drs_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	drsCmd.AddCommand(drs_listTagsForResourceCmd)
}
