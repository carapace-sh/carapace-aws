package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var healthlake_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Add a user-specifed key and value tag to a data store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(healthlake_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(healthlake_tagResourceCmd).Standalone()

		healthlake_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) that grants access to the data store tags are being added to.")
		healthlake_tagResourceCmd.Flags().String("tags", "", "The user-specified key and value pair tags being added to a data store.")
		healthlake_tagResourceCmd.MarkFlagRequired("resource-arn")
		healthlake_tagResourceCmd.MarkFlagRequired("tags")
	})
	healthlakeCmd.AddCommand(healthlake_tagResourceCmd)
}
