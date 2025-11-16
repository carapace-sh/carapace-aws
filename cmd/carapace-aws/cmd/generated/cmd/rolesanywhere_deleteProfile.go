package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolesanywhere_deleteProfileCmd = &cobra.Command{
	Use:   "delete-profile",
	Short: "Deletes a profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolesanywhere_deleteProfileCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rolesanywhere_deleteProfileCmd).Standalone()

		rolesanywhere_deleteProfileCmd.Flags().String("profile-id", "", "The unique identifier of the profile.")
		rolesanywhere_deleteProfileCmd.MarkFlagRequired("profile-id")
	})
	rolesanywhereCmd.AddCommand(rolesanywhere_deleteProfileCmd)
}
