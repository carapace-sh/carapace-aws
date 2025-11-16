package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53profiles_createProfileCmd = &cobra.Command{
	Use:   "create-profile",
	Short: "Creates an empty Route 53 Profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53profiles_createProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(route53profiles_createProfileCmd).Standalone()

		route53profiles_createProfileCmd.Flags().String("client-token", "", "`ClientToken` is an idempotency token that ensures a call to `CreateProfile` completes only once.")
		route53profiles_createProfileCmd.Flags().String("name", "", "A name for the Profile.")
		route53profiles_createProfileCmd.Flags().String("tags", "", "A list of the tag keys and values that you want to associate with the Route 53 Profile.")
		route53profiles_createProfileCmd.MarkFlagRequired("client-token")
		route53profiles_createProfileCmd.MarkFlagRequired("name")
	})
	route53profilesCmd.AddCommand(route53profiles_createProfileCmd)
}
