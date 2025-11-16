package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var dynamodb_describeLimitsCmd = &cobra.Command{
	Use:   "describe-limits",
	Short: "Returns the current provisioned-capacity quotas for your Amazon Web Services account in a Region, both for the Region as a whole and for any one DynamoDB table that you create there.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(dynamodb_describeLimitsCmd).Standalone()

	dynamodbCmd.AddCommand(dynamodb_describeLimitsCmd)
}
