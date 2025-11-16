package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_createProtectionGroupCmd = &cobra.Command{
	Use:   "create-protection-group",
	Short: "Creates a grouping of protected resources so they can be handled as a collective.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_createProtectionGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(shield_createProtectionGroupCmd).Standalone()

		shield_createProtectionGroupCmd.Flags().String("aggregation", "", "Defines how Shield combines resource data for the group in order to detect, mitigate, and report events.")
		shield_createProtectionGroupCmd.Flags().String("members", "", "The Amazon Resource Names (ARNs) of the resources to include in the protection group.")
		shield_createProtectionGroupCmd.Flags().String("pattern", "", "The criteria to use to choose the protected resources for inclusion in the group.")
		shield_createProtectionGroupCmd.Flags().String("protection-group-id", "", "The name of the protection group.")
		shield_createProtectionGroupCmd.Flags().String("resource-type", "", "The resource type to include in the protection group.")
		shield_createProtectionGroupCmd.Flags().String("tags", "", "One or more tag key-value pairs for the protection group.")
		shield_createProtectionGroupCmd.MarkFlagRequired("aggregation")
		shield_createProtectionGroupCmd.MarkFlagRequired("pattern")
		shield_createProtectionGroupCmd.MarkFlagRequired("protection-group-id")
	})
	shieldCmd.AddCommand(shield_createProtectionGroupCmd)
}
