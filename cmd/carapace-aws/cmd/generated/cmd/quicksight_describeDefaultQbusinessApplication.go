package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeDefaultQbusinessApplicationCmd = &cobra.Command{
	Use:   "describe-default-qbusiness-application",
	Short: "Describes a Amazon Q Business application that is linked to an Quick Sight account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeDefaultQbusinessApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_describeDefaultQbusinessApplicationCmd).Standalone()

		quicksight_describeDefaultQbusinessApplicationCmd.Flags().String("aws-account-id", "", "The ID of the Quick Sight account that is linked to the Amazon Q Business application that you want described.")
		quicksight_describeDefaultQbusinessApplicationCmd.Flags().String("namespace", "", "The Quick Sight namespace that contains the linked Amazon Q Business application.")
		quicksight_describeDefaultQbusinessApplicationCmd.MarkFlagRequired("aws-account-id")
	})
	quicksightCmd.AddCommand(quicksight_describeDefaultQbusinessApplicationCmd)
}
