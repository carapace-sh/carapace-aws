package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var datazone_acceptPredictionsCmd = &cobra.Command{
	Use:   "accept-predictions",
	Short: "Accepts automatically generated business-friendly metadata for your Amazon DataZone assets.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(datazone_acceptPredictionsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(datazone_acceptPredictionsCmd).Standalone()

		datazone_acceptPredictionsCmd.Flags().String("accept-choices", "", "Specifies the prediction (aka, the automatically generated piece of metadata) and the target (for example, a column name) that can be accepted.")
		datazone_acceptPredictionsCmd.Flags().String("accept-rule", "", "Specifies the rule (or the conditions) under which a prediction can be accepted.")
		datazone_acceptPredictionsCmd.Flags().String("client-token", "", "A unique, case-sensitive identifier to ensure idempotency of the request.")
		datazone_acceptPredictionsCmd.Flags().String("domain-identifier", "", "The identifier of the Amazon DataZone domain.")
		datazone_acceptPredictionsCmd.Flags().String("identifier", "", "The identifier of the asset.")
		datazone_acceptPredictionsCmd.Flags().String("revision", "", "The revision that is to be made to the asset.")
		datazone_acceptPredictionsCmd.MarkFlagRequired("domain-identifier")
		datazone_acceptPredictionsCmd.MarkFlagRequired("identifier")
	})
	datazoneCmd.AddCommand(datazone_acceptPredictionsCmd)
}
