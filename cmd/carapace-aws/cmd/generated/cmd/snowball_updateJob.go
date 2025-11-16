package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var snowball_updateJobCmd = &cobra.Command{
	Use:   "update-job",
	Short: "While a job's `JobState` value is `New`, you can update some of the information associated with a job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(snowball_updateJobCmd).Standalone()

	snowball_updateJobCmd.Flags().String("address-id", "", "The ID of the updated [Address]() object.")
	snowball_updateJobCmd.Flags().String("description", "", "The updated description of this job's [JobMetadata]() object.")
	snowball_updateJobCmd.Flags().String("forwarding-address-id", "", "The updated ID for the forwarding address for a job.")
	snowball_updateJobCmd.Flags().String("job-id", "", "The job ID of the job that you want to update, for example `JID123e4567-e89b-12d3-a456-426655440000`.")
	snowball_updateJobCmd.Flags().String("notification", "", "The new or updated [Notification]() object.")
	snowball_updateJobCmd.Flags().String("on-device-service-configuration", "", "Specifies the service or services on the Snow Family device that your transferred data will be exported from or imported into.")
	snowball_updateJobCmd.Flags().String("pickup-details", "", "")
	snowball_updateJobCmd.Flags().String("resources", "", "The updated `JobResource` object, or the updated [JobResource]() object.")
	snowball_updateJobCmd.Flags().String("role-arn", "", "The new role Amazon Resource Name (ARN) that you want to associate with this job.")
	snowball_updateJobCmd.Flags().String("shipping-option", "", "The updated shipping option value of this job's [ShippingDetails]() object.")
	snowball_updateJobCmd.Flags().String("snowball-capacity-preference", "", "The updated `SnowballCapacityPreference` of this job's [JobMetadata]() object.")
	snowball_updateJobCmd.MarkFlagRequired("job-id")
	snowballCmd.AddCommand(snowball_updateJobCmd)
}
