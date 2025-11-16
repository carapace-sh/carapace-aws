package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var shield_describeAttackStatisticsCmd = &cobra.Command{
	Use:   "describe-attack-statistics",
	Short: "Provides information about the number and type of attacks Shield has detected in the last year for all resources that belong to your account, regardless of whether you've defined Shield protections for them.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(shield_describeAttackStatisticsCmd).Standalone()

	shieldCmd.AddCommand(shield_describeAttackStatisticsCmd)
}
