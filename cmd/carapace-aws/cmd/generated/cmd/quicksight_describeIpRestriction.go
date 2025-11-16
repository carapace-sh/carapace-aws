package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeIpRestrictionCmd = &cobra.Command{
	Use:   "describe-ip-restriction",
	Short: "Provides a summary and status of IP rules.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeIpRestrictionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_describeIpRestrictionCmd).Standalone()

		quicksight_describeIpRestrictionCmd.Flags().String("aws-account-id", "", "The ID of the Amazon Web Services account that contains the IP rules.")
		quicksight_describeIpRestrictionCmd.MarkFlagRequired("aws-account-id")
	})
	quicksightCmd.AddCommand(quicksight_describeIpRestrictionCmd)
}
