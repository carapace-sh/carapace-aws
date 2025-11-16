package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_startUpdateSignalMapCmd = &cobra.Command{
	Use:   "start-update-signal-map",
	Short: "Initiates an update for the specified signal map.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_startUpdateSignalMapCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(medialive_startUpdateSignalMapCmd).Standalone()

		medialive_startUpdateSignalMapCmd.Flags().String("cloud-watch-alarm-template-group-identifiers", "", "")
		medialive_startUpdateSignalMapCmd.Flags().String("description", "", "A resource's optional description.")
		medialive_startUpdateSignalMapCmd.Flags().String("discovery-entry-point-arn", "", "A top-level supported AWS resource ARN to discovery a signal map from.")
		medialive_startUpdateSignalMapCmd.Flags().String("event-bridge-rule-template-group-identifiers", "", "")
		medialive_startUpdateSignalMapCmd.Flags().String("force-rediscovery", "", "If true, will force a rediscovery of a signal map if an unchanged discoveryEntryPointArn is provided.")
		medialive_startUpdateSignalMapCmd.Flags().String("identifier", "", "A signal map's identifier.")
		medialive_startUpdateSignalMapCmd.Flags().String("name", "", "A resource's name.")
		medialive_startUpdateSignalMapCmd.MarkFlagRequired("identifier")
	})
	medialiveCmd.AddCommand(medialive_startUpdateSignalMapCmd)
}
