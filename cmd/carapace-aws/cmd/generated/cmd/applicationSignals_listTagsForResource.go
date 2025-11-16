package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationSignals_listTagsForResourceCmd = &cobra.Command{
	Use:   "list-tags-for-resource",
	Short: "Displays the tags associated with a CloudWatch resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationSignals_listTagsForResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationSignals_listTagsForResourceCmd).Standalone()

		applicationSignals_listTagsForResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the CloudWatch resource that you want to view tags for.")
		applicationSignals_listTagsForResourceCmd.MarkFlagRequired("resource-arn")
	})
	applicationSignalsCmd.AddCommand(applicationSignals_listTagsForResourceCmd)
}
