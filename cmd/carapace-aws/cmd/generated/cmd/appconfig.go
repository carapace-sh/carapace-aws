package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var appconfigCmd = &cobra.Command{
	Use:   "appconfig",
	Short: "AppConfig feature flags and dynamic configurations help software builders quickly and securely adjust application behavior in production environments without full code deployments.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(appconfigCmd).Standalone()

	rootCmd.AddCommand(appconfigCmd)
}
