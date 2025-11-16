package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_stopUserImportJobCmd = &cobra.Command{
	Use:   "stop-user-import-job",
	Short: "Instructs your user pool to stop a running job that's importing users from a CSV file that contains their usernames and attributes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_stopUserImportJobCmd).Standalone()

	cognitoIdp_stopUserImportJobCmd.Flags().String("job-id", "", "The ID of a running user import job.")
	cognitoIdp_stopUserImportJobCmd.Flags().String("user-pool-id", "", "The ID of the user pool that you want to stop.")
	cognitoIdp_stopUserImportJobCmd.MarkFlagRequired("job-id")
	cognitoIdp_stopUserImportJobCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_stopUserImportJobCmd)
}
