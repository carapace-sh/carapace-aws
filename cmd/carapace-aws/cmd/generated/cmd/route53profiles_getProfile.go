package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53profiles_getProfileCmd = &cobra.Command{
	Use:   "get-profile",
	Short: "Returns information about a specified Route 53 Profile, such as whether whether the Profile is shared, and the current status of the Profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53profiles_getProfileCmd).Standalone()

	route53profiles_getProfileCmd.Flags().String("profile-id", "", "ID of the Profile.")
	route53profiles_getProfileCmd.MarkFlagRequired("profile-id")
	route53profilesCmd.AddCommand(route53profiles_getProfileCmd)
}
