package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_updateThreatEntitySetCmd = &cobra.Command{
	Use:   "update-threat-entity-set",
	Short: "Updates the threat entity set associated with the specified `threatEntitySetId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_updateThreatEntitySetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_updateThreatEntitySetCmd).Standalone()

		guardduty_updateThreatEntitySetCmd.Flags().Bool("activate", false, "A boolean value that indicates whether GuardDuty is to start using this updated threat entity set.")
		guardduty_updateThreatEntitySetCmd.Flags().String("detector-id", "", "The unique ID of the GuardDuty detector associated with the threat entity set that you want to update.")
		guardduty_updateThreatEntitySetCmd.Flags().String("expected-bucket-owner", "", "The Amazon Web Services account ID that owns the Amazon S3 bucket specified in the **location** parameter.")
		guardduty_updateThreatEntitySetCmd.Flags().String("location", "", "The URI of the file that contains the trusted entity set.")
		guardduty_updateThreatEntitySetCmd.Flags().String("name", "", "A user-friendly name to identify the trusted entity set.")
		guardduty_updateThreatEntitySetCmd.Flags().Bool("no-activate", false, "A boolean value that indicates whether GuardDuty is to start using this updated threat entity set.")
		guardduty_updateThreatEntitySetCmd.Flags().String("threat-entity-set-id", "", "The ID returned by GuardDuty after updating the threat entity set resource.")
		guardduty_updateThreatEntitySetCmd.MarkFlagRequired("detector-id")
		guardduty_updateThreatEntitySetCmd.Flag("no-activate").Hidden = true
		guardduty_updateThreatEntitySetCmd.MarkFlagRequired("threat-entity-set-id")
	})
	guarddutyCmd.AddCommand(guardduty_updateThreatEntitySetCmd)
}
