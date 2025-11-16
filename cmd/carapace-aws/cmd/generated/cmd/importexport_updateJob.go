package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var importexport_updateJobCmd = &cobra.Command{
	Use:   "update-job",
	Short: "You use this operation to change the parameters specified in the original manifest file by supplying a new manifest file.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(importexport_updateJobCmd).Standalone()

	importexport_updateJobCmd.Flags().String("apiversion", "", "")
	importexport_updateJobCmd.Flags().String("job-id", "", "")
	importexport_updateJobCmd.Flags().String("job-type", "", "")
	importexport_updateJobCmd.Flags().String("manifest", "", "")
	importexport_updateJobCmd.Flags().String("validate-only", "", "")
	importexport_updateJobCmd.MarkFlagRequired("job-id")
	importexport_updateJobCmd.MarkFlagRequired("job-type")
	importexport_updateJobCmd.MarkFlagRequired("manifest")
	importexport_updateJobCmd.MarkFlagRequired("validate-only")
	importexportCmd.AddCommand(importexport_updateJobCmd)
}
