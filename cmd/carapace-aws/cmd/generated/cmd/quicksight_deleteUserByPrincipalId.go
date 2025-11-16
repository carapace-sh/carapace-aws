package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var quicksight_deleteUserByPrincipalIdCmd = &cobra.Command{
	Use:   "delete-user-by-principal-id",
	Short: "Deletes a user identified by its principal ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(quicksight_deleteUserByPrincipalIdCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(quicksight_deleteUserByPrincipalIdCmd).Standalone()

		quicksight_deleteUserByPrincipalIdCmd.Flags().String("aws-account-id", "", "The ID for the Amazon Web Services account that the user is in.")
		quicksight_deleteUserByPrincipalIdCmd.Flags().String("namespace", "", "The namespace.")
		quicksight_deleteUserByPrincipalIdCmd.Flags().String("principal-id", "", "The principal ID of the user.")
		quicksight_deleteUserByPrincipalIdCmd.MarkFlagRequired("aws-account-id")
		quicksight_deleteUserByPrincipalIdCmd.MarkFlagRequired("namespace")
		quicksight_deleteUserByPrincipalIdCmd.MarkFlagRequired("principal-id")
	})
	quicksightCmd.AddCommand(quicksight_deleteUserByPrincipalIdCmd)
}
