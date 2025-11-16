package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var medialive_createInputSecurityGroupCmd = &cobra.Command{
	Use:   "create-input-security-group",
	Short: "Creates a Input Security Group",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(medialive_createInputSecurityGroupCmd).Standalone()

	medialive_createInputSecurityGroupCmd.Flags().String("tags", "", "A collection of key-value pairs.")
	medialive_createInputSecurityGroupCmd.Flags().String("whitelist-rules", "", "List of IPv4 CIDR addresses to whitelist")
	medialiveCmd.AddCommand(medialive_createInputSecurityGroupCmd)
}
