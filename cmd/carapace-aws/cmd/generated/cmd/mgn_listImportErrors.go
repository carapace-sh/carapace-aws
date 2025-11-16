package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mgn_listImportErrorsCmd = &cobra.Command{
	Use:   "list-import-errors",
	Short: "List import errors.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mgn_listImportErrorsCmd).Standalone()

	mgn_listImportErrorsCmd.Flags().String("import-id", "", "List import errors request import id.")
	mgn_listImportErrorsCmd.Flags().String("max-results", "", "List import errors request max results.")
	mgn_listImportErrorsCmd.Flags().String("next-token", "", "List import errors request next token.")
	mgn_listImportErrorsCmd.MarkFlagRequired("import-id")
	mgnCmd.AddCommand(mgn_listImportErrorsCmd)
}
