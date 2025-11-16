package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdentity_describeIdentityPoolCmd = &cobra.Command{
	Use:   "describe-identity-pool",
	Short: "Gets details about a particular identity pool, including the pool name, ID description, creation date, and current number of users.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdentity_describeIdentityPoolCmd).Standalone()

	cognitoIdentity_describeIdentityPoolCmd.Flags().String("identity-pool-id", "", "An identity pool ID in the format REGION:GUID.")
	cognitoIdentity_describeIdentityPoolCmd.MarkFlagRequired("identity-pool-id")
	cognitoIdentityCmd.AddCommand(cognitoIdentity_describeIdentityPoolCmd)
}
