package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_createAuthenticationProfileCmd = &cobra.Command{
	Use:   "create-authentication-profile",
	Short: "Creates an authentication profile with the specified parameters.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_createAuthenticationProfileCmd).Standalone()

	redshift_createAuthenticationProfileCmd.Flags().String("authentication-profile-content", "", "The content of the authentication profile in JSON format.")
	redshift_createAuthenticationProfileCmd.Flags().String("authentication-profile-name", "", "The name of the authentication profile to be created.")
	redshift_createAuthenticationProfileCmd.MarkFlagRequired("authentication-profile-content")
	redshift_createAuthenticationProfileCmd.MarkFlagRequired("authentication-profile-name")
	redshiftCmd.AddCommand(redshift_createAuthenticationProfileCmd)
}
