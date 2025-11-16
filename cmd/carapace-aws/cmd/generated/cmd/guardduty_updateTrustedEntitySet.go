package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_updateTrustedEntitySetCmd = &cobra.Command{
	Use:   "update-trusted-entity-set",
	Short: "Updates the trusted entity set associated with the specified `trustedEntitySetId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_updateTrustedEntitySetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_updateTrustedEntitySetCmd).Standalone()

		guardduty_updateTrustedEntitySetCmd.Flags().Bool("activate", false, "A boolean value that indicates whether GuardDuty is to start using this updated trusted entity set.")
		guardduty_updateTrustedEntitySetCmd.Flags().String("detector-id", "", "The unique ID of the GuardDuty detector associated with the threat entity set that you want to update.")
		guardduty_updateTrustedEntitySetCmd.Flags().String("expected-bucket-owner", "", "The Amazon Web Services account ID that owns the Amazon S3 bucket specified in the **location** parameter.")
		guardduty_updateTrustedEntitySetCmd.Flags().String("location", "", "The URI of the file that contains the trusted entity set.")
		guardduty_updateTrustedEntitySetCmd.Flags().String("name", "", "A user-friendly name to identify the trusted entity set.")
		guardduty_updateTrustedEntitySetCmd.Flags().Bool("no-activate", false, "A boolean value that indicates whether GuardDuty is to start using this updated trusted entity set.")
		guardduty_updateTrustedEntitySetCmd.Flags().String("trusted-entity-set-id", "", "The ID returned by GuardDuty after updating the trusted entity set resource.")
		guardduty_updateTrustedEntitySetCmd.MarkFlagRequired("detector-id")
		guardduty_updateTrustedEntitySetCmd.Flag("no-activate").Hidden = true
		guardduty_updateTrustedEntitySetCmd.MarkFlagRequired("trusted-entity-set-id")
	})
	guarddutyCmd.AddCommand(guardduty_updateTrustedEntitySetCmd)
}
