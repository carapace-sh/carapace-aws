package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var healthlakeCmd = &cobra.Command{
	Use:   "healthlake",
	Short: "This is the *AWS HealthLake API Reference*.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(healthlakeCmd).Standalone()

	rootCmd.AddCommand(healthlakeCmd)
}
