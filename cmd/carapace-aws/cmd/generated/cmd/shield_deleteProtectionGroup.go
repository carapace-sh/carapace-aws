package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_deleteProtectionGroupCmd = &cobra.Command{
	Use:   "delete-protection-group",
	Short: "Removes the specified protection group.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_deleteProtectionGroupCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(shield_deleteProtectionGroupCmd).Standalone()

		shield_deleteProtectionGroupCmd.Flags().String("protection-group-id", "", "The name of the protection group.")
		shield_deleteProtectionGroupCmd.MarkFlagRequired("protection-group-id")
	})
	shieldCmd.AddCommand(shield_deleteProtectionGroupCmd)
}
