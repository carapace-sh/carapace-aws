package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var location_createTrackerCmd = &cobra.Command{
	Use:   "create-tracker",
	Short: "Creates a tracker resource in your Amazon Web Services account, which lets you retrieve current and historical location of devices.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(location_createTrackerCmd).Standalone()

	location_createTrackerCmd.Flags().String("description", "", "An optional description for the tracker resource.")
	location_createTrackerCmd.Flags().Bool("event-bridge-enabled", false, "Whether to enable position `UPDATE` events from this tracker to be sent to EventBridge.")
	location_createTrackerCmd.Flags().Bool("kms-key-enable-geospatial-queries", false, "Enables `GeospatialQueries` for a tracker that uses a [Amazon Web Services KMS customer managed key](https://docs.aws.amazon.com/kms/latest/developerguide/create-keys.html).")
	location_createTrackerCmd.Flags().String("kms-key-id", "", "A key identifier for an [Amazon Web Services KMS customer managed key](https://docs.aws.amazon.com/kms/latest/developerguide/create-keys.html).")
	location_createTrackerCmd.Flags().Bool("no-event-bridge-enabled", false, "Whether to enable position `UPDATE` events from this tracker to be sent to EventBridge.")
	location_createTrackerCmd.Flags().Bool("no-kms-key-enable-geospatial-queries", false, "Enables `GeospatialQueries` for a tracker that uses a [Amazon Web Services KMS customer managed key](https://docs.aws.amazon.com/kms/latest/developerguide/create-keys.html).")
	location_createTrackerCmd.Flags().String("position-filtering", "", "Specifies the position filtering for the tracker resource.")
	location_createTrackerCmd.Flags().String("pricing-plan", "", "No longer used.")
	location_createTrackerCmd.Flags().String("pricing-plan-data-source", "", "This parameter is no longer used.")
	location_createTrackerCmd.Flags().String("tags", "", "Applies one or more tags to the tracker resource.")
	location_createTrackerCmd.Flags().String("tracker-name", "", "The name for the tracker resource.")
	location_createTrackerCmd.Flag("no-event-bridge-enabled").Hidden = true
	location_createTrackerCmd.Flag("no-kms-key-enable-geospatial-queries").Hidden = true
	location_createTrackerCmd.MarkFlagRequired("tracker-name")
	locationCmd.AddCommand(location_createTrackerCmd)
}
