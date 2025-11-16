package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var launchWizardCmd = &cobra.Command{
	Use:   "launch-wizard",
	Short: "Launch Wizard offers a guided way of sizing, configuring, and deploying Amazon Web Services resources for third party applications, such as Microsoft SQL Server Always On and HANA based SAP systems, without the need to manually identify and provision individual Amazon Web Services resources.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(launchWizardCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(launchWizardCmd).Standalone()

	})
	rootCmd.AddCommand(launchWizardCmd)
}
