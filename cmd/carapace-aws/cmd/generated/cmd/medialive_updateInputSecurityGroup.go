package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_updateInputSecurityGroupCmd = &cobra.Command{
	Use:   "update-input-security-group",
	Short: "Update an Input Security Group's Whilelists.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_updateInputSecurityGroupCmd).Standalone()

	medialive_updateInputSecurityGroupCmd.Flags().String("input-security-group-id", "", "The id of the Input Security Group to update.")
	medialive_updateInputSecurityGroupCmd.Flags().String("tags", "", "A collection of key-value pairs.")
	medialive_updateInputSecurityGroupCmd.Flags().String("whitelist-rules", "", "List of IPv4 CIDR addresses to whitelist")
	medialive_updateInputSecurityGroupCmd.MarkFlagRequired("input-security-group-id")
	medialiveCmd.AddCommand(medialive_updateInputSecurityGroupCmd)
}
