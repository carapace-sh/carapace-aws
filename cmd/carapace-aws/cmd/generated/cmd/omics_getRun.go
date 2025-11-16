package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_getRunCmd = &cobra.Command{
	Use:   "get-run",
	Short: "Gets detailed information about a specific run using its ID.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_getRunCmd).Standalone()

	omics_getRunCmd.Flags().String("export", "", "The run's export format.")
	omics_getRunCmd.Flags().String("id", "", "The run's ID.")
	omics_getRunCmd.MarkFlagRequired("id")
	omicsCmd.AddCommand(omics_getRunCmd)
}
