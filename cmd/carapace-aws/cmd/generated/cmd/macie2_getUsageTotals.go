package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var macie2_getUsageTotalsCmd = &cobra.Command{
	Use:   "get-usage-totals",
	Short: "Retrieves (queries) aggregated usage data for an account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(macie2_getUsageTotalsCmd).Standalone()

	macie2_getUsageTotalsCmd.Flags().String("time-range", "", "The inclusive time period to retrieve the data for.")
	macie2Cmd.AddCommand(macie2_getUsageTotalsCmd)
}
