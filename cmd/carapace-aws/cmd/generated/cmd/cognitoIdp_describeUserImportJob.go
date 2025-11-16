package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_describeUserImportJobCmd = &cobra.Command{
	Use:   "describe-user-import-job",
	Short: "Describes a user import job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_describeUserImportJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdp_describeUserImportJobCmd).Standalone()

		cognitoIdp_describeUserImportJobCmd.Flags().String("job-id", "", "The Id of the user import job that you want to describe.")
		cognitoIdp_describeUserImportJobCmd.Flags().String("user-pool-id", "", "The ID of the user pool that's associated with the import job.")
		cognitoIdp_describeUserImportJobCmd.MarkFlagRequired("job-id")
		cognitoIdp_describeUserImportJobCmd.MarkFlagRequired("user-pool-id")
	})
	cognitoIdpCmd.AddCommand(cognitoIdp_describeUserImportJobCmd)
}
