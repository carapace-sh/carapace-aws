package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_getTrustedEntitySetCmd = &cobra.Command{
	Use:   "get-trusted-entity-set",
	Short: "Retrieves the trusted entity set associated with the specified `trustedEntitySetId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_getTrustedEntitySetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(guardduty_getTrustedEntitySetCmd).Standalone()

		guardduty_getTrustedEntitySetCmd.Flags().String("detector-id", "", "The unique ID of the GuardDuty detector associated with this trusted entity set.")
		guardduty_getTrustedEntitySetCmd.Flags().String("trusted-entity-set-id", "", "The unique ID that helps GuardDuty identify the trusted entity set.")
		guardduty_getTrustedEntitySetCmd.MarkFlagRequired("detector-id")
		guardduty_getTrustedEntitySetCmd.MarkFlagRequired("trusted-entity-set-id")
	})
	guarddutyCmd.AddCommand(guardduty_getTrustedEntitySetCmd)
}
