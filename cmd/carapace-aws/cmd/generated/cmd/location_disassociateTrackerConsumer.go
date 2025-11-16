package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_disassociateTrackerConsumerCmd = &cobra.Command{
	Use:   "disassociate-tracker-consumer",
	Short: "Removes the association between a tracker resource and a geofence collection.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_disassociateTrackerConsumerCmd).Standalone()

	location_disassociateTrackerConsumerCmd.Flags().String("consumer-arn", "", "The Amazon Resource Name (ARN) for the geofence collection to be disassociated from the tracker resource.")
	location_disassociateTrackerConsumerCmd.Flags().String("tracker-name", "", "The name of the tracker resource to be dissociated from the consumer.")
	location_disassociateTrackerConsumerCmd.MarkFlagRequired("consumer-arn")
	location_disassociateTrackerConsumerCmd.MarkFlagRequired("tracker-name")
	locationCmd.AddCommand(location_disassociateTrackerConsumerCmd)
}
