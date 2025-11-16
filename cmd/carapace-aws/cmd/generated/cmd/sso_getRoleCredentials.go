package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sso_getRoleCredentialsCmd = &cobra.Command{
	Use:   "get-role-credentials",
	Short: "Returns the STS short-term credentials for a given role name that is assigned to the user.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sso_getRoleCredentialsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sso_getRoleCredentialsCmd).Standalone()

		sso_getRoleCredentialsCmd.Flags().String("access-token", "", "The token issued by the `CreateToken` API call.")
		sso_getRoleCredentialsCmd.Flags().String("account-id", "", "The identifier for the AWS account that is assigned to the user.")
		sso_getRoleCredentialsCmd.Flags().String("role-name", "", "The friendly name of the role that is assigned to the user.")
		sso_getRoleCredentialsCmd.MarkFlagRequired("access-token")
		sso_getRoleCredentialsCmd.MarkFlagRequired("account-id")
		sso_getRoleCredentialsCmd.MarkFlagRequired("role-name")
	})
	ssoCmd.AddCommand(sso_getRoleCredentialsCmd)
}
