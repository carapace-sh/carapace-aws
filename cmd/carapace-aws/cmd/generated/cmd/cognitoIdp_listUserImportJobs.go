package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_listUserImportJobsCmd = &cobra.Command{
	Use:   "list-user-import-jobs",
	Short: "Given a user pool ID, returns user import jobs and their details.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_listUserImportJobsCmd).Standalone()

	cognitoIdp_listUserImportJobsCmd.Flags().String("max-results", "", "The maximum number of import jobs that you want Amazon Cognito to return in the response.")
	cognitoIdp_listUserImportJobsCmd.Flags().String("pagination-token", "", "This API operation returns a limited number of results.")
	cognitoIdp_listUserImportJobsCmd.Flags().String("user-pool-id", "", "The ID of the user pool where you want to list import jobs.")
	cognitoIdp_listUserImportJobsCmd.MarkFlagRequired("max-results")
	cognitoIdp_listUserImportJobsCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_listUserImportJobsCmd)
}
