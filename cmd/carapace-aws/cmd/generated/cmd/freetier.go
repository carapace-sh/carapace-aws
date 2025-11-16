package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var freetierCmd = &cobra.Command{
	Use:   "freetier",
	Short: "You can use the Amazon Web Services Free Tier API to query programmatically your Free Tier usage data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(freetierCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(freetierCmd).Standalone()

	})
	rootCmd.AddCommand(freetierCmd)
}
