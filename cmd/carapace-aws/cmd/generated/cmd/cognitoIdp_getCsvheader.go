package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_getCsvheaderCmd = &cobra.Command{
	Use:   "get-csvheader",
	Short: "Given a user pool ID, generates a comma-separated value (CSV) list populated with available user attributes in the user pool.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_getCsvheaderCmd).Standalone()

	cognitoIdp_getCsvheaderCmd.Flags().String("user-pool-id", "", "The ID of the user pool that you want to import users into.")
	cognitoIdp_getCsvheaderCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_getCsvheaderCmd)
}
