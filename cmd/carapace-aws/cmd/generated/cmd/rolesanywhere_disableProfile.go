package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolesanywhere_disableProfileCmd = &cobra.Command{
	Use:   "disable-profile",
	Short: "Disables a profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolesanywhere_disableProfileCmd).Standalone()

	rolesanywhere_disableProfileCmd.Flags().String("profile-id", "", "The unique identifier of the profile.")
	rolesanywhere_disableProfileCmd.MarkFlagRequired("profile-id")
	rolesanywhereCmd.AddCommand(rolesanywhere_disableProfileCmd)
}
