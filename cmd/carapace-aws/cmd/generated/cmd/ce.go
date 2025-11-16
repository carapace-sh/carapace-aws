package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var ceCmd = &cobra.Command{
	Use:   "ce",
	Short: "You can use the Cost Explorer API to programmatically query your cost and usage data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(ceCmd).Standalone()

	rootCmd.AddCommand(ceCmd)
}
