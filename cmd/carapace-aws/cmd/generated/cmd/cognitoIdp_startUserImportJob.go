package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_startUserImportJobCmd = &cobra.Command{
	Use:   "start-user-import-job",
	Short: "Instructs your user pool to start importing users from a CSV file that contains their usernames and attributes.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_startUserImportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_startUserImportJobCmd).Standalone()

		cognitoIdp_startUserImportJobCmd.Flags().String("job-id", "", "The ID of a user import job that you previously created.")
		cognitoIdp_startUserImportJobCmd.Flags().String("user-pool-id", "", "The ID of the user pool that you want to start importing users into.")
		cognitoIdp_startUserImportJobCmd.MarkFlagRequired("job-id")
		cognitoIdp_startUserImportJobCmd.MarkFlagRequired("user-pool-id")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_startUserImportJobCmd)
}
