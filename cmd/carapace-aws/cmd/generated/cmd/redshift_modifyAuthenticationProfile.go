package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_modifyAuthenticationProfileCmd = &cobra.Command{
	Use:   "modify-authentication-profile",
	Short: "Modifies an authentication profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_modifyAuthenticationProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(redshift_modifyAuthenticationProfileCmd).Standalone()

		redshift_modifyAuthenticationProfileCmd.Flags().String("authentication-profile-content", "", "The new content of the authentication profile in JSON format.")
		redshift_modifyAuthenticationProfileCmd.Flags().String("authentication-profile-name", "", "The name of the authentication profile to replace.")
		redshift_modifyAuthenticationProfileCmd.MarkFlagRequired("authentication-profile-content")
		redshift_modifyAuthenticationProfileCmd.MarkFlagRequired("authentication-profile-name")
	})
	redshiftCmd.AddCommand(redshift_modifyAuthenticationProfileCmd)
}
