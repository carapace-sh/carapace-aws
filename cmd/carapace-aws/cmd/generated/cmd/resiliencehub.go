package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var resiliencehubCmd = &cobra.Command{
	Use:   "resiliencehub",
	Short: "Resilience Hub helps you proactively prepare and protect your Amazon Web Services applications from disruptions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(resiliencehubCmd).Standalone()

	rootCmd.AddCommand(resiliencehubCmd)
}
