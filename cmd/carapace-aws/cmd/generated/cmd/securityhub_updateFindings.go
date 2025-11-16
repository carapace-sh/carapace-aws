package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var securityhub_updateFindingsCmd = &cobra.Command{
	Use:   "update-findings",
	Short: "`UpdateFindings` is a deprecated operation.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(securityhub_updateFindingsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(securityhub_updateFindingsCmd).Standalone()

		securityhub_updateFindingsCmd.Flags().String("filters", "", "A collection of attributes that specify which findings you want to update.")
		securityhub_updateFindingsCmd.Flags().String("note", "", "The updated note for the finding.")
		securityhub_updateFindingsCmd.Flags().String("record-state", "", "The updated record state for the finding.")
		securityhub_updateFindingsCmd.MarkFlagRequired("filters")
	})
	securityhubCmd.AddCommand(securityhub_updateFindingsCmd)
}
