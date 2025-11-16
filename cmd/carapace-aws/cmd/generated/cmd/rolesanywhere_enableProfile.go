package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rolesanywhere_enableProfileCmd = &cobra.Command{
	Use:   "enable-profile",
	Short: "Enables temporary credential requests for a profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rolesanywhere_enableProfileCmd).Standalone()

	rolesanywhere_enableProfileCmd.Flags().String("profile-id", "", "The unique identifier of the profile.")
	rolesanywhere_enableProfileCmd.MarkFlagRequired("profile-id")
	rolesanywhereCmd.AddCommand(rolesanywhere_enableProfileCmd)
}
