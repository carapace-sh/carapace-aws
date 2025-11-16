package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudtrail_createEventDataStoreCmd = &cobra.Command{
	Use:   "create-event-data-store",
	Short: "Creates a new event data store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudtrail_createEventDataStoreCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudtrail_createEventDataStoreCmd).Standalone()

		cloudtrail_createEventDataStoreCmd.Flags().String("advanced-event-selectors", "", "The advanced event selectors to use to select the events for the data store.")
		cloudtrail_createEventDataStoreCmd.Flags().String("billing-mode", "", "The billing mode for the event data store determines the cost for ingesting events and the default and maximum retention period for the event data store.")
		cloudtrail_createEventDataStoreCmd.Flags().String("kms-key-id", "", "Specifies the KMS key ID to use to encrypt the events delivered by CloudTrail.")
		cloudtrail_createEventDataStoreCmd.Flags().Bool("multi-region-enabled", false, "Specifies whether the event data store includes events from all Regions, or only from the Region in which the event data store is created.")
		cloudtrail_createEventDataStoreCmd.Flags().String("name", "", "The name of the event data store.")
		cloudtrail_createEventDataStoreCmd.Flags().Bool("no-multi-region-enabled", false, "Specifies whether the event data store includes events from all Regions, or only from the Region in which the event data store is created.")
		cloudtrail_createEventDataStoreCmd.Flags().Bool("no-organization-enabled", false, "Specifies whether an event data store collects events logged for an organization in Organizations.")
		cloudtrail_createEventDataStoreCmd.Flags().Bool("no-start-ingestion", false, "Specifies whether the event data store should start ingesting live events.")
		cloudtrail_createEventDataStoreCmd.Flags().Bool("organization-enabled", false, "Specifies whether an event data store collects events logged for an organization in Organizations.")
		cloudtrail_createEventDataStoreCmd.Flags().String("retention-period", "", "The retention period of the event data store, in days.")
		cloudtrail_createEventDataStoreCmd.Flags().Bool("start-ingestion", false, "Specifies whether the event data store should start ingesting live events.")
		cloudtrail_createEventDataStoreCmd.Flags().String("tags-list", "", "")
		cloudtrail_createEventDataStoreCmd.Flags().String("termination-protection-enabled", "", "Specifies whether termination protection is enabled for the event data store.")
		cloudtrail_createEventDataStoreCmd.MarkFlagRequired("name")
		cloudtrail_createEventDataStoreCmd.Flag("no-multi-region-enabled").Hidden = true
		cloudtrail_createEventDataStoreCmd.Flag("no-organization-enabled").Hidden = true
		cloudtrail_createEventDataStoreCmd.Flag("no-start-ingestion").Hidden = true
	})
	cloudtrailCmd.AddCommand(cloudtrail_createEventDataStoreCmd)
}
