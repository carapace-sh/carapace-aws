package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emrContainers_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns tags to resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emrContainers_tagResourceCmd).Standalone()

	emrContainers_tagResourceCmd.Flags().String("resource-arn", "", "The ARN of resources.")
	emrContainers_tagResourceCmd.Flags().String("tags", "", "The tags assigned to resources.")
	emrContainers_tagResourceCmd.MarkFlagRequired("resource-arn")
	emrContainers_tagResourceCmd.MarkFlagRequired("tags")
	emrContainersCmd.AddCommand(emrContainers_tagResourceCmd)
}
