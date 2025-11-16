package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Adds or overwrites only the specified tags for the specified Application Migration Service resource or resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_tagResourceCmd).Standalone()

		mgn_tagResourceCmd.Flags().String("resource-arn", "", "Tag resource by ARN.")
		mgn_tagResourceCmd.Flags().String("tags", "", "Tag resource by Tags.")
		mgn_tagResourceCmd.MarkFlagRequired("resource-arn")
		mgn_tagResourceCmd.MarkFlagRequired("tags")
	})
	mgnCmd.AddCommand(mgn_tagResourceCmd)
}
