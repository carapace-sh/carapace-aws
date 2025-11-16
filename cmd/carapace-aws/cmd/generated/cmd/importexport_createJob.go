package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var importexport_createJobCmd = &cobra.Command{
	Use:   "create-job",
	Short: "This operation initiates the process of scheduling an upload or download of your data.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importexport_createJobCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(importexport_createJobCmd).Standalone()

		importexport_createJobCmd.Flags().String("apiversion", "", "")
		importexport_createJobCmd.Flags().String("job-type", "", "")
		importexport_createJobCmd.Flags().String("manifest", "", "")
		importexport_createJobCmd.Flags().String("manifest-addendum", "", "")
		importexport_createJobCmd.Flags().String("validate-only", "", "")
		importexport_createJobCmd.MarkFlagRequired("job-type")
		importexport_createJobCmd.MarkFlagRequired("manifest")
		importexport_createJobCmd.MarkFlagRequired("validate-only")
	})
	importexportCmd.AddCommand(importexport_createJobCmd)
}
