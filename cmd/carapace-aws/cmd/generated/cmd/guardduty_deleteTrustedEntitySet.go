package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_deleteTrustedEntitySetCmd = &cobra.Command{
	Use:   "delete-trusted-entity-set",
	Short: "Deletes the trusted entity set that is associated with the specified `trustedEntitySetId`.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_deleteTrustedEntitySetCmd).Standalone()

	guardduty_deleteTrustedEntitySetCmd.Flags().String("detector-id", "", "The unique ID of the detector associated with the trusted entity set resource.")
	guardduty_deleteTrustedEntitySetCmd.Flags().String("trusted-entity-set-id", "", "The unique ID that helps GuardDuty identify which trusted entity set needs to be deleted.")
	guardduty_deleteTrustedEntitySetCmd.MarkFlagRequired("detector-id")
	guardduty_deleteTrustedEntitySetCmd.MarkFlagRequired("trusted-entity-set-id")
	guarddutyCmd.AddCommand(guardduty_deleteTrustedEntitySetCmd)
}
