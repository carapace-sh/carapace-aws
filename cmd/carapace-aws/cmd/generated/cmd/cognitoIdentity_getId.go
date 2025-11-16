package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdentity_getIdCmd = &cobra.Command{
	Use:   "get-id",
	Short: "Generates (or retrieves) IdentityID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdentity_getIdCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cognitoIdentity_getIdCmd).Standalone()

		cognitoIdentity_getIdCmd.Flags().String("account-id", "", "A standard Amazon Web Services account ID (9+ digits).")
		cognitoIdentity_getIdCmd.Flags().String("identity-pool-id", "", "An identity pool ID in the format REGION:GUID.")
		cognitoIdentity_getIdCmd.Flags().String("logins", "", "A set of optional name-value pairs that map provider names to provider tokens.")
		cognitoIdentity_getIdCmd.MarkFlagRequired("identity-pool-id")
	})
	cognitoIdentityCmd.AddCommand(cognitoIdentity_getIdCmd)
}
