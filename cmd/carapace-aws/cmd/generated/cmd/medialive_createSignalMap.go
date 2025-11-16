package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_createSignalMapCmd = &cobra.Command{
	Use:   "create-signal-map",
	Short: "Initiates the creation of a new signal map.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_createSignalMapCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_createSignalMapCmd).Standalone()

		medialive_createSignalMapCmd.Flags().String("cloud-watch-alarm-template-group-identifiers", "", "")
		medialive_createSignalMapCmd.Flags().String("description", "", "A resource's optional description.")
		medialive_createSignalMapCmd.Flags().String("discovery-entry-point-arn", "", "A top-level supported AWS resource ARN to discovery a signal map from.")
		medialive_createSignalMapCmd.Flags().String("event-bridge-rule-template-group-identifiers", "", "")
		medialive_createSignalMapCmd.Flags().String("name", "", "A resource's name.")
		medialive_createSignalMapCmd.Flags().String("request-id", "", "An ID that you assign to a create request.")
		medialive_createSignalMapCmd.Flags().String("tags", "", "")
		medialive_createSignalMapCmd.MarkFlagRequired("discovery-entry-point-arn")
		medialive_createSignalMapCmd.MarkFlagRequired("name")
	})
	medialiveCmd.AddCommand(medialive_createSignalMapCmd)
}
