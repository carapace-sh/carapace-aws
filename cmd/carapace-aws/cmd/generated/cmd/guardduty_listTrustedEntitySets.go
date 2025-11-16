package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var guardduty_listTrustedEntitySetsCmd = &cobra.Command{
	Use:   "list-trusted-entity-sets",
	Short: "Lists the trusted entity sets associated with the specified GuardDuty detector ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(guardduty_listTrustedEntitySetsCmd).Standalone()

	guardduty_listTrustedEntitySetsCmd.Flags().String("detector-id", "", "The unique ID of the GuardDuty detector that is associated with this threat entity set.")
	guardduty_listTrustedEntitySetsCmd.Flags().String("max-results", "", "You can use this parameter to indicate the maximum number of items you want in the response.")
	guardduty_listTrustedEntitySetsCmd.Flags().String("next-token", "", "You can use this parameter when paginating results.")
	guardduty_listTrustedEntitySetsCmd.MarkFlagRequired("detector-id")
	guarddutyCmd.AddCommand(guardduty_listTrustedEntitySetsCmd)
}
