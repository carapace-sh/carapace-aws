package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_listExportErrorsCmd = &cobra.Command{
	Use:   "list-export-errors",
	Short: "List export errors.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_listExportErrorsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mgn_listExportErrorsCmd).Standalone()

		mgn_listExportErrorsCmd.Flags().String("export-id", "", "List export errors request export id.")
		mgn_listExportErrorsCmd.Flags().String("max-results", "", "List export errors request max results.")
		mgn_listExportErrorsCmd.Flags().String("next-token", "", "List export errors request next token.")
		mgn_listExportErrorsCmd.MarkFlagRequired("export-id")
	})
	mgnCmd.AddCommand(mgn_listExportErrorsCmd)
}
