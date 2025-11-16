package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_rejectPredictionsCmd = &cobra.Command{
	Use:   "reject-predictions",
	Short: "Rejects automatically generated business-friendly metadata for your Amazon DataZone assets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_rejectPredictionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_rejectPredictionsCmd).Standalone()

		datazone_rejectPredictionsCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier that is provided to ensure the idempotency of the request.")
		datazone_rejectPredictionsCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain.")
		datazone_rejectPredictionsCmd.Flags().String("identifier", "", "The identifier of the prediction.")
		datazone_rejectPredictionsCmd.Flags().String("reject-choices", "", "Specifies the prediction (aka, the automatically generated piece of metadata) and the target (for example, a column name) that can be rejected.")
		datazone_rejectPredictionsCmd.Flags().String("reject-rule", "", "Specifies the rule (or the conditions) under which a prediction can be rejected.")
		datazone_rejectPredictionsCmd.Flags().String("revision", "", "The revision that is to be made to the asset.")
		datazone_rejectPredictionsCmd.MarkFlagRequired("domain-identifier")
		datazone_rejectPredictionsCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_rejectPredictionsCmd)
}
