package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "List all tags for your Application Migration Service resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_listTagsForResourceCmd).Standalone()

		mgn_listTagsForResourceCmd.Flags().String("resource-arn", "", "List tags for resource request by ARN.")
		mgn_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	mgnCmd.AddCommand(mgn_listTagsForResourceCmd)
}
