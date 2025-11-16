package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_describeUserPoolCmd = &cobra.Command{
	Use:   "describe-user-pool",
	Short: "Given a user pool ID, returns configuration information.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_describeUserPoolCmd).Standalone()

	cognitoIdp_describeUserPoolCmd.Flags().String("user-pool-id", "", "The ID of the user pool you want to describe.")
	cognitoIdp_describeUserPoolCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_describeUserPoolCmd)
}
