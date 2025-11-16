package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_deleteSecurityProfileCmd = &cobra.Command{
	Use:   "delete-security-profile",
	Short: "Deletes a security profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_deleteSecurityProfileCmd).Standalone()

	connect_deleteSecurityProfileCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_deleteSecurityProfileCmd.Flags().String("security-profile-id", "", "The identifier for the security profle.")
	connect_deleteSecurityProfileCmd.MarkFlagRequired("instance-id")
	connect_deleteSecurityProfileCmd.MarkFlagRequired("security-profile-id")
	connectCmd.AddCommand(connect_deleteSecurityProfileCmd)
}
