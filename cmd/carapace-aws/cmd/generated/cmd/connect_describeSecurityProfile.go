package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_describeSecurityProfileCmd = &cobra.Command{
	Use:   "describe-security-profile",
	Short: "Gets basic information about the security profile.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_describeSecurityProfileCmd).Standalone()

	connect_describeSecurityProfileCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
	connect_describeSecurityProfileCmd.Flags().String("security-profile-id", "", "The identifier for the security profle.")
	connect_describeSecurityProfileCmd.MarkFlagRequired("instance-id")
	connect_describeSecurityProfileCmd.MarkFlagRequired("security-profile-id")
	connectCmd.AddCommand(connect_describeSecurityProfileCmd)
}
