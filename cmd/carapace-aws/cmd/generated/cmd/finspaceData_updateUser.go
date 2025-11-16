package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_updateUserCmd = &cobra.Command{
	Use:   "update-user",
	Short: "Modifies the details of the specified user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_updateUserCmd).Standalone()

	finspaceData_updateUserCmd.Flags().String("api-access", "", "The option to indicate whether the user can use the `GetProgrammaticAccessCredentials` API to obtain credentials that can then be used to access other FinSpace Data API operations.")
	finspaceData_updateUserCmd.Flags().String("api-access-principal-arn", "", "The ARN identifier of an AWS user or role that is allowed to call the `GetProgrammaticAccessCredentials` API to obtain a credentials token for a specific FinSpace user.")
	finspaceData_updateUserCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
	finspaceData_updateUserCmd.Flags().String("first-name", "", "The first name of the user.")
	finspaceData_updateUserCmd.Flags().String("last-name", "", "The last name of the user.")
	finspaceData_updateUserCmd.Flags().String("type", "", "The option to indicate the type of user.")
	finspaceData_updateUserCmd.Flags().String("user-id", "", "The unique identifier for the user that you want to update.")
	finspaceData_updateUserCmd.MarkFlagRequired("user-id")
	finspaceDataCmd.AddCommand(finspaceData_updateUserCmd)
}
