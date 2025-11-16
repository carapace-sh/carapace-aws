package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var importexport_cancelJobCmd = &cobra.Command{
	Use:   "cancel-job",
	Short: "This operation cancels a specified job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importexport_cancelJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(importexport_cancelJobCmd).Standalone()

		importexport_cancelJobCmd.Flags().String("apiversion", "", "")
		importexport_cancelJobCmd.Flags().String("job-id", "", "")
		importexport_cancelJobCmd.MarkFlagRequired("job-id")
	})
	importexportCmd.AddCommand(importexport_cancelJobCmd)
}
