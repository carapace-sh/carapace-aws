package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoSync_listIdentityPoolUsageCmd = &cobra.Command{
	Use:   "list-identity-pool-usage",
	Short: "Gets a list of identity pools registered with Cognito.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoSync_listIdentityPoolUsageCmd).Standalone()

	cognitoSync_listIdentityPoolUsageCmd.Flags().String("max-results", "", "The maximum number of results to be returned.")
	cognitoSync_listIdentityPoolUsageCmd.Flags().String("next-token", "", "A pagination token for obtaining the next page of results.")
	cognitoSyncCmd.AddCommand(cognitoSync_listIdentityPoolUsageCmd)
}
