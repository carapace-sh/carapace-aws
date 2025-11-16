package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var drs_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds or overwrites only the specified tags for the specified Elastic Disaster Recovery resource or resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(drs_tagResourceCmd).Standalone()

	drs_tagResourceCmd.Flags().String("resource-arn", "", "ARN of the resource for which tags are to be added or updated.")
	drs_tagResourceCmd.Flags().String("tags", "", "Array of tags to be added or updated.")
	drs_tagResourceCmd.MarkFlagRequired("resource-arn")
	drs_tagResourceCmd.MarkFlagRequired("tags")
	drsCmd.AddCommand(drs_tagResourceCmd)
}
