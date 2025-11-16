package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var observabilityadminCmd = &cobra.Command{
	Use:   "observabilityadmin",
	Short: "You can use Amazon CloudWatch Observability Admin to discover and understand the state of telemetry configuration in CloudWatch for your Amazon Web Services Organization or account.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(observabilityadminCmd).Standalone()

	rootCmd.AddCommand(observabilityadminCmd)
}
