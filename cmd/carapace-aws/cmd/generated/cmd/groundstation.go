package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var groundstationCmd = &cobra.Command{
	Use:   "groundstation",
	Short: "Welcome to the AWS Ground Station API Reference.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(groundstationCmd).Standalone()

	rootCmd.AddCommand(groundstationCmd)
}
