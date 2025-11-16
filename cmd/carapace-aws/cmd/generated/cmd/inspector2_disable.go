package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var inspector2_disableCmd = &cobra.Command{
	Use:   "disable",
	Short: "Disables Amazon Inspector scans for one or more Amazon Web Services accounts.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(inspector2_disableCmd).Standalone()

	inspector2_disableCmd.Flags().String("account-ids", "", "An array of account IDs you want to disable Amazon Inspector scans for.")
	inspector2_disableCmd.Flags().String("resource-types", "", "The resource scan types you want to disable.")
	inspector2Cmd.AddCommand(inspector2_disableCmd)
}
