package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_batchGetFreeTrialInfoCmd = &cobra.Command{
	Use:   "batch-get-free-trial-info",
	Short: "Gets free trial status for multiple Amazon Web Services accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_batchGetFreeTrialInfoCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(inspector2_batchGetFreeTrialInfoCmd).Standalone()

		inspector2_batchGetFreeTrialInfoCmd.Flags().String("account-ids", "", "The account IDs to get free trial status for.")
		inspector2_batchGetFreeTrialInfoCmd.MarkFlagRequired("account-ids")
	})
	inspector2Cmd.AddCommand(inspector2_batchGetFreeTrialInfoCmd)
}
