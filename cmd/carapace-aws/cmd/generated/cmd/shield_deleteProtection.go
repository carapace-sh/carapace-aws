package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_deleteProtectionCmd = &cobra.Command{
	Use:   "delete-protection",
	Short: "Deletes an Shield Advanced [Protection]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_deleteProtectionCmd).Standalone()

	shield_deleteProtectionCmd.Flags().String("protection-id", "", "The unique identifier (ID) for the [Protection]() object to be deleted.")
	shield_deleteProtectionCmd.MarkFlagRequired("protection-id")
	shieldCmd.AddCommand(shield_deleteProtectionCmd)
}
