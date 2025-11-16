package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var healthlake_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Returns a list of all existing tags associated with a data store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(healthlake_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(healthlake_listTagsForResourceCmd).Standalone()

		healthlake_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the data store to which tags are being added.")
		healthlake_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	healthlakeCmd.AddCommand(healthlake_listTagsForResourceCmd)
}
