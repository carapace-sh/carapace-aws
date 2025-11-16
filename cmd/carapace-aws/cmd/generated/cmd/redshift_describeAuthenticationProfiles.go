package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var redshift_describeAuthenticationProfilesCmd = &cobra.Command{
	Use:   "describe-authentication-profiles",
	Short: "Describes an authentication profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(redshift_describeAuthenticationProfilesCmd).Standalone()

	redshift_describeAuthenticationProfilesCmd.Flags().String("authentication-profile-name", "", "The name of the authentication profile to describe.")
	redshiftCmd.AddCommand(redshift_describeAuthenticationProfilesCmd)
}
