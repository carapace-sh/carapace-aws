package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_createUserImportJobCmd = &cobra.Command{
	Use:   "create-user-import-job",
	Short: "Creates a user import job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_createUserImportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_createUserImportJobCmd).Standalone()

		cognitoIdp_createUserImportJobCmd.Flags().String("cloud-watch-logs-role-arn", "", "You must specify an IAM role that has permission to log import-job results to Amazon CloudWatch Logs.")
		cognitoIdp_createUserImportJobCmd.Flags().String("job-name", "", "A friendly name for the user import job.")
		cognitoIdp_createUserImportJobCmd.Flags().String("user-pool-id", "", "The ID of the user pool that you want to import users into.")
		cognitoIdp_createUserImportJobCmd.MarkFlagRequired("cloud-watch-logs-role-arn")
		cognitoIdp_createUserImportJobCmd.MarkFlagRequired("job-name")
		cognitoIdp_createUserImportJobCmd.MarkFlagRequired("user-pool-id")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_createUserImportJobCmd)
}
