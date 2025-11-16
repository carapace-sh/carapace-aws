package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_updateEventDataStoreCmd = &cobra.Command{
	Use:   "update-event-data-store",
	Short: "Updates an event data store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_updateEventDataStoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudtrail_updateEventDataStoreCmd).Standalone()

		cloudtrail_updateEventDataStoreCmd.Flags().String("advanced-event-selectors", "", "The advanced event selectors used to select events for the event data store.")
		cloudtrail_updateEventDataStoreCmd.Flags().String("billing-mode", "", "You can't change the billing mode from `EXTENDABLE_RETENTION_PRICING` to `FIXED_RETENTION_PRICING`.")
		cloudtrail_updateEventDataStoreCmd.Flags().String("event-data-store", "", "The ARN (or the ID suffix of the ARN) of the event data store that you want to update.")
		cloudtrail_updateEventDataStoreCmd.Flags().String("kms-key-id", "", "Specifies the KMS key ID to use to encrypt the events delivered by CloudTrail.")
		cloudtrail_updateEventDataStoreCmd.Flags().Bool("multi-region-enabled", false, "Specifies whether an event data store collects events from all Regions, or only from the Region in which it was created.")
		cloudtrail_updateEventDataStoreCmd.Flags().String("name", "", "The event data store name.")
		cloudtrail_updateEventDataStoreCmd.Flags().Bool("no-multi-region-enabled", false, "Specifies whether an event data store collects events from all Regions, or only from the Region in which it was created.")
		cloudtrail_updateEventDataStoreCmd.Flags().Bool("no-organization-enabled", false, "Specifies whether an event data store collects events logged for an organization in Organizations.")
		cloudtrail_updateEventDataStoreCmd.Flags().Bool("organization-enabled", false, "Specifies whether an event data store collects events logged for an organization in Organizations.")
		cloudtrail_updateEventDataStoreCmd.Flags().String("retention-period", "", "The retention period of the event data store, in days.")
		cloudtrail_updateEventDataStoreCmd.Flags().String("termination-protection-enabled", "", "Indicates that termination protection is enabled and the event data store cannot be automatically deleted.")
		cloudtrail_updateEventDataStoreCmd.MarkFlagRequired("event-data-store")
		cloudtrail_updateEventDataStoreCmd.Flag("no-multi-region-enabled").Hidden = true
		cloudtrail_updateEventDataStoreCmd.Flag("no-organization-enabled").Hidden = true
	})
	cloudtrailCmd.AddCommand(cloudtrail_updateEventDataStoreCmd)
}
