package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_updateProtectionGroupCmd = &cobra.Command{
	Use:   "update-protection-group",
	Short: "Updates an existing protection group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_updateProtectionGroupCmd).Standalone()

	shield_updateProtectionGroupCmd.Flags().String("aggregation", "", "Defines how Shield combines resource data for the group in order to detect, mitigate, and report events.")
	shield_updateProtectionGroupCmd.Flags().String("members", "", "The Amazon Resource Names (ARNs) of the resources to include in the protection group.")
	shield_updateProtectionGroupCmd.Flags().String("pattern", "", "The criteria to use to choose the protected resources for inclusion in the group.")
	shield_updateProtectionGroupCmd.Flags().String("protection-group-id", "", "The name of the protection group.")
	shield_updateProtectionGroupCmd.Flags().String("resource-type", "", "The resource type to include in the protection group.")
	shield_updateProtectionGroupCmd.MarkFlagRequired("aggregation")
	shield_updateProtectionGroupCmd.MarkFlagRequired("pattern")
	shield_updateProtectionGroupCmd.MarkFlagRequired("protection-group-id")
	shieldCmd.AddCommand(shield_updateProtectionGroupCmd)
}
