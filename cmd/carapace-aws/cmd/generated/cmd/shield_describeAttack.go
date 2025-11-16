package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_describeAttackCmd = &cobra.Command{
	Use:   "describe-attack",
	Short: "Describes the details of a DDoS attack.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_describeAttackCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(shield_describeAttackCmd).Standalone()

		shield_describeAttackCmd.Flags().String("attack-id", "", "The unique identifier (ID) for the attack.")
		shield_describeAttackCmd.MarkFlagRequired("attack-id")
	})
	shieldCmd.AddCommand(shield_describeAttackCmd)
}
