package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lookoutequipment_listRetrainingSchedulersCmd = &cobra.Command{
	Use:   "list-retraining-schedulers",
	Short: "Lists all retraining schedulers in your account, filtering by model name prefix and status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lookoutequipment_listRetrainingSchedulersCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lookoutequipment_listRetrainingSchedulersCmd).Standalone()

		lookoutequipment_listRetrainingSchedulersCmd.Flags().String("max-results", "", "Specifies the maximum number of retraining schedulers to list.")
		lookoutequipment_listRetrainingSchedulersCmd.Flags().String("model-name-begins-with", "", "Specify this field to only list retraining schedulers whose machine learning models begin with the value you specify.")
		lookoutequipment_listRetrainingSchedulersCmd.Flags().String("next-token", "", "If the number of results exceeds the maximum, a pagination token is returned.")
		lookoutequipment_listRetrainingSchedulersCmd.Flags().String("status", "", "Specify this field to only list retraining schedulers whose status matches the value you specify.")
	})
	lookoutequipmentCmd.AddCommand(lookoutequipment_listRetrainingSchedulersCmd)
}
