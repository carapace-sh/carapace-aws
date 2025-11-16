package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var billingconductorCmd = &cobra.Command{
	Use:   "billingconductor",
	Short: "Billing Conductor is a fully managed service that you can use to customize a [pro forma](https://docs.aws.amazon.com/billingconductor/latest/userguide/understanding-eb.html#eb-other-definitions) version of your billing data each month, to accurately show or chargeback your end customers.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(billingconductorCmd).Standalone()

	rootCmd.AddCommand(billingconductorCmd)
}
