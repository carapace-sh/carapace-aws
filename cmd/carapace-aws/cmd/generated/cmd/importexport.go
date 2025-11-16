package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var importexportCmd = &cobra.Command{
	Use:   "importexport",
	Short: "AWS Import/Export Service AWS Import/Export accelerates transferring large amounts of data between the AWS cloud and portable storage devices that you mail to us.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importexportCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(importexportCmd).Standalone()

	})
	rootCmd.AddCommand(importexportCmd)
}
