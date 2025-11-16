package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_listRefreshSchedulesCmd = &cobra.Command{
	Use:   "list-refresh-schedules",
	Short: "Lists the refresh schedules of a dataset.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_listRefreshSchedulesCmd).Standalone()

	quicksight_listRefreshSchedulesCmd.Flags().String("aws-account-id", "", "The Amazon Web Services account ID.")
	quicksight_listRefreshSchedulesCmd.Flags().String("data-set-id", "", "The ID of the dataset.")
	quicksight_listRefreshSchedulesCmd.MarkFlagRequired("aws-account-id")
	quicksight_listRefreshSchedulesCmd.MarkFlagRequired("data-set-id")
	quicksightCmd.AddCommand(quicksight_listRefreshSchedulesCmd)
}
