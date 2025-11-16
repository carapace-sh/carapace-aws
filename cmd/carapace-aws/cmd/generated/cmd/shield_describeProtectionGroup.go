package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_describeProtectionGroupCmd = &cobra.Command{
	Use:   "describe-protection-group",
	Short: "Returns the specification for the specified protection group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_describeProtectionGroupCmd).Standalone()

	shield_describeProtectionGroupCmd.Flags().String("protection-group-id", "", "The name of the protection group.")
	shield_describeProtectionGroupCmd.MarkFlagRequired("protection-group-id")
	shieldCmd.AddCommand(shield_describeProtectionGroupCmd)
}
