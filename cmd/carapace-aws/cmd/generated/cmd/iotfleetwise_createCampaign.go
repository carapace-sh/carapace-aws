package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotfleetwise_createCampaignCmd = &cobra.Command{
	Use:   "create-campaign",
	Short: "Creates an orchestration of data collection rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotfleetwise_createCampaignCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(iotfleetwise_createCampaignCmd).Standalone()

		iotfleetwise_createCampaignCmd.Flags().String("collection-scheme", "", "The data collection scheme associated with the campaign.")
		iotfleetwise_createCampaignCmd.Flags().String("compression", "", "Determines whether to compress signals before transmitting data to Amazon Web Services IoT FleetWise.")
		iotfleetwise_createCampaignCmd.Flags().String("data-destination-configs", "", "The destination where the campaign sends data.")
		iotfleetwise_createCampaignCmd.Flags().String("data-extra-dimensions", "", "A list of vehicle attributes to associate with a campaign.")
		iotfleetwise_createCampaignCmd.Flags().String("data-partitions", "", "The data partitions associated with the signals collected from the vehicle.")
		iotfleetwise_createCampaignCmd.Flags().String("description", "", "An optional description of the campaign to help identify its purpose.")
		iotfleetwise_createCampaignCmd.Flags().String("diagnostics-mode", "", "Option for a vehicle to send diagnostic trouble codes to Amazon Web Services IoT FleetWise.")
		iotfleetwise_createCampaignCmd.Flags().String("expiry-time", "", "The time the campaign expires, in seconds since epoch (January 1, 1970 at midnight UTC time).")
		iotfleetwise_createCampaignCmd.Flags().String("name", "", "The name of the campaign to create.")
		iotfleetwise_createCampaignCmd.Flags().String("post-trigger-collection-duration", "", "How long (in milliseconds) to collect raw data after a triggering event initiates the collection.")
		iotfleetwise_createCampaignCmd.Flags().String("priority", "", "A number indicating the priority of one campaign over another campaign for a certain vehicle or fleet.")
		iotfleetwise_createCampaignCmd.Flags().String("signal-catalog-arn", "", "The Amazon Resource Name (ARN) of the signal catalog to associate with the campaign.")
		iotfleetwise_createCampaignCmd.Flags().String("signals-to-collect", "", "A list of information about signals to collect.")
		iotfleetwise_createCampaignCmd.Flags().String("signals-to-fetch", "", "A list of information about signals to fetch.")
		iotfleetwise_createCampaignCmd.Flags().String("spooling-mode", "", "Determines whether to store collected data after a vehicle lost a connection with the cloud.")
		iotfleetwise_createCampaignCmd.Flags().String("start-time", "", "The time, in milliseconds, to deliver a campaign after it was approved.")
		iotfleetwise_createCampaignCmd.Flags().String("tags", "", "Metadata that can be used to manage the campaign.")
		iotfleetwise_createCampaignCmd.Flags().String("target-arn", "", "The ARN of the vehicle or fleet to deploy a campaign to.")
		iotfleetwise_createCampaignCmd.MarkFlagRequired("collection-scheme")
		iotfleetwise_createCampaignCmd.MarkFlagRequired("name")
		iotfleetwise_createCampaignCmd.MarkFlagRequired("signal-catalog-arn")
		iotfleetwise_createCampaignCmd.MarkFlagRequired("target-arn")
	})
	iotfleetwiseCmd.AddCommand(iotfleetwise_createCampaignCmd)
}
