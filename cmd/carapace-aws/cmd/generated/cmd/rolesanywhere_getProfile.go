package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolesanywhere_getProfileCmd = &cobra.Command{
	Use:   "get-profile",
	Short: "Gets a profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolesanywhere_getProfileCmd).Standalone()

	rolesanywhere_getProfileCmd.Flags().String("profile-id", "", "The unique identifier of the profile.")
	rolesanywhere_getProfileCmd.MarkFlagRequired("profile-id")
	rolesanywhereCmd.AddCommand(rolesanywhere_getProfileCmd)
}
