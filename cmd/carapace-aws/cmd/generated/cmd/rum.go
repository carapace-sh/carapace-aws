package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rumCmd = &cobra.Command{
	Use:   "rum",
	Short: "With Amazon CloudWatch RUM, you can perform real-user monitoring to collect client-side data about your web application performance from actual user sessions in real time.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rumCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(rumCmd).Standalone()

	})
	rootCmd.AddCommand(rumCmd)
}
