package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_createUserCmd = &cobra.Command{
	Use:   "create-user",
	Short: "Creates a new user in FinSpace.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_createUserCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(finspaceData_createUserCmd).Standalone()

		finspaceData_createUserCmd.Flags().String("api-access", "", "The option to indicate whether the user can use the `GetProgrammaticAccessCredentials` API to obtain credentials that can then be used to access other FinSpace Data API operations.")
		finspaceData_createUserCmd.Flags().String("api-access-principal-arn", "", "The ARN identifier of an AWS user or role that is allowed to call the `GetProgrammaticAccessCredentials` API to obtain a credentials token for a specific FinSpace user.")
		finspaceData_createUserCmd.Flags().String("client-token", "", "A token that ensures idempotency.")
		finspaceData_createUserCmd.Flags().String("email-address", "", "The email address of the user that you want to register.")
		finspaceData_createUserCmd.Flags().String("first-name", "", "The first name of the user that you want to register.")
		finspaceData_createUserCmd.Flags().String("last-name", "", "The last name of the user that you want to register.")
		finspaceData_createUserCmd.Flags().String("type", "", "The option to indicate the type of user.")
		finspaceData_createUserCmd.MarkFlagRequired("email-address")
		finspaceData_createUserCmd.MarkFlagRequired("type")
	})
	finspaceDataCmd.AddCommand(finspaceData_createUserCmd)
}
