package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteUserCmd = &cobra.Command{
	Use:   "delete-user",
	Short: "Deletes the Amazon Quick Sight user that is associated with the identity of the IAM user or role that's making the call.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteUserCmd).Standalone()

	quicksight_deleteUserCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that the user is in.")
	quicksight_deleteUserCmd.Flags().String("namespace", "", "The namespace.")
	quicksight_deleteUserCmd.Flags().String("user-name", "", "The name of the user that you want to delete.")
	quicksight_deleteUserCmd.MarkFlagRequired("aws-account-id")
	quicksight_deleteUserCmd.MarkFlagRequired("namespace")
	quicksight_deleteUserCmd.MarkFlagRequired("user-name")
	quicksightCmd.AddCommand(quicksight_deleteUserCmd)
}
