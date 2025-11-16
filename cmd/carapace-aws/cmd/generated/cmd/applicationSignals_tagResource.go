package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var applicationSignals_tagResourceCmd = &cobra.Command{
	Use:   "tag-resource",
	Short: "Assigns one or more tags (key-value pairs) to the specified CloudWatch resource, such as a service level objective.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(applicationSignals_tagResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(applicationSignals_tagResourceCmd).Standalone()

		applicationSignals_tagResourceCmd.Flags().String("resource-arn", "", "The Amazon Resource Name (ARN) of the CloudWatch resource that you want to set tags for.")
		applicationSignals_tagResourceCmd.Flags().String("tags", "", "The list of key-value pairs to associate with the alarm.")
		applicationSignals_tagResourceCmd.MarkFlagRequired("resource-arn")
		applicationSignals_tagResourceCmd.MarkFlagRequired("tags")
	})
	applicationSignalsCmd.AddCommand(applicationSignals_tagResourceCmd)
}
