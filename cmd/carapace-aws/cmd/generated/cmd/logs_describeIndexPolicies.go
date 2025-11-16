package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var logs_describeIndexPoliciesCmd = &cobra.Command{
	Use:   "describe-index-policies",
	Short: "Returns the field index policies of the specified log group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(logs_describeIndexPoliciesCmd).Standalone()

	logs_describeIndexPoliciesCmd.Flags().String("log-group-identifiers", "", "An array containing the name or ARN of the log group that you want to retrieve field index policies for.")
	logs_describeIndexPoliciesCmd.Flags().String("next-token", "", "")
	logs_describeIndexPoliciesCmd.MarkFlagRequired("log-group-identifiers")
	logsCmd.AddCommand(logs_describeIndexPoliciesCmd)
}
