package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_describeTimeToLiveCmd = &cobra.Command{
	Use:   "describe-time-to-live",
	Short: "Gives a description of the Time to Live (TTL) status on the specified table.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_describeTimeToLiveCmd).Standalone()

	dynamodb_describeTimeToLiveCmd.Flags().String("table-name", "", "The name of the table to be described.")
	dynamodb_describeTimeToLiveCmd.MarkFlagRequired("table-name")
	dynamodbCmd.AddCommand(dynamodb_describeTimeToLiveCmd)
}
