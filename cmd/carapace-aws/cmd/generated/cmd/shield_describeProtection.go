package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_describeProtectionCmd = &cobra.Command{
	Use:   "describe-protection",
	Short: "Lists the details of a [Protection]() object.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_describeProtectionCmd).Standalone()

	shield_describeProtectionCmd.Flags().String("protection-id", "", "The unique identifier (ID) for the [Protection]() object to describe.")
	shield_describeProtectionCmd.Flags().String("resource-arn", "", "The ARN (Amazon Resource Name) of the protected Amazon Web Services resource.")
	shieldCmd.AddCommand(shield_describeProtectionCmd)
}
