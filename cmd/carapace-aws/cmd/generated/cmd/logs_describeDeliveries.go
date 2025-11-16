package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_describeDeliveriesCmd = &cobra.Command{
	Use:   "describe-deliveries",
	Short: "Retrieves a list of the deliveries that have been created in the account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_describeDeliveriesCmd).Standalone()

	logs_describeDeliveriesCmd.Flags().String("limit", "", "Optionally specify the maximum number of deliveries to return in the response.")
	logs_describeDeliveriesCmd.Flags().String("next-token", "", "")
	logsCmd.AddCommand(logs_describeDeliveriesCmd)
}
