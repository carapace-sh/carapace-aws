package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_associateTrackerConsumerCmd = &cobra.Command{
	Use:   "associate-tracker-consumer",
	Short: "Creates an association between a geofence collection and a tracker resource.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_associateTrackerConsumerCmd).Standalone()

	location_associateTrackerConsumerCmd.Flags().String("consumer-arn", "", "The Amazon Resource Name (ARN) for the geofence collection to be associated to tracker resource.")
	location_associateTrackerConsumerCmd.Flags().String("tracker-name", "", "The name of the tracker resource to be associated with a geofence collection.")
	location_associateTrackerConsumerCmd.MarkFlagRequired("consumer-arn")
	location_associateTrackerConsumerCmd.MarkFlagRequired("tracker-name")
	locationCmd.AddCommand(location_associateTrackerConsumerCmd)
}
