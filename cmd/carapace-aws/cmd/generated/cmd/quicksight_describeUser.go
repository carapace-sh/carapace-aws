package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_describeUserCmd = &cobra.Command{
	Use:   "describe-user",
	Short: "Returns information about a user, given the user name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_describeUserCmd).Standalone()

	quicksight_describeUserCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that the user is in.")
	quicksight_describeUserCmd.Flags().String("namespace", "", "The namespace.")
	quicksight_describeUserCmd.Flags().String("user-name", "", "The name of the user that you want to describe.")
	quicksight_describeUserCmd.MarkFlagRequired("aws-account-id")
	quicksight_describeUserCmd.MarkFlagRequired("namespace")
	quicksight_describeUserCmd.MarkFlagRequired("user-name")
	quicksightCmd.AddCommand(quicksight_describeUserCmd)
}
