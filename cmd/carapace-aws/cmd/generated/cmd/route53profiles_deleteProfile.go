package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var route53profiles_deleteProfileCmd = &cobra.Command{
	Use:   "delete-profile",
	Short: "Deletes the specified Route 53 Profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(route53profiles_deleteProfileCmd).Standalone()

	route53profiles_deleteProfileCmd.Flags().String("profile-id", "", "The ID of the Profile that you want to delete.")
	route53profiles_deleteProfileCmd.MarkFlagRequired("profile-id")
	route53profilesCmd.AddCommand(route53profiles_deleteProfileCmd)
}
