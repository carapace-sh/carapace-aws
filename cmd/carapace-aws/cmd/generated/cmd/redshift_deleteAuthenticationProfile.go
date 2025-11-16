package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_deleteAuthenticationProfileCmd = &cobra.Command{
	Use:   "delete-authentication-profile",
	Short: "Deletes an authentication profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_deleteAuthenticationProfileCmd).Standalone()

	redshift_deleteAuthenticationProfileCmd.Flags().String("authentication-profile-name", "", "The name of the authentication profile to delete.")
	redshift_deleteAuthenticationProfileCmd.MarkFlagRequired("authentication-profile-name")
	redshiftCmd.AddCommand(redshift_deleteAuthenticationProfileCmd)
}
