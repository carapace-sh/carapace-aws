package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cognitoIdp_describeResourceServerCmd = &cobra.Command{
	Use:   "describe-resource-server",
	Short: "Describes a resource server.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cognitoIdp_describeResourceServerCmd).Standalone()

	cognitoIdp_describeResourceServerCmd.Flags().String("identifier", "", "A unique resource server identifier for the resource server.")
	cognitoIdp_describeResourceServerCmd.Flags().String("user-pool-id", "", "The ID of the user pool that hosts the resource server.")
	cognitoIdp_describeResourceServerCmd.MarkFlagRequired("identifier")
	cognitoIdp_describeResourceServerCmd.MarkFlagRequired("user-pool-id")
	cognitoIdpCmd.AddCommand(cognitoIdp_describeResourceServerCmd)
}
