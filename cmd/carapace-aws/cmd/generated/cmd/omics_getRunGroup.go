package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var omics_getRunGroupCmd = &cobra.Command{
	Use:   "get-run-group",
	Short: "Gets information about a run group and returns its metadata.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(omics_getRunGroupCmd).Standalone()

	omics_getRunGroupCmd.Flags().String("id", "", "The group's ID.")
	omics_getRunGroupCmd.MarkFlagRequired("id")
	omicsCmd.AddCommand(omics_getRunGroupCmd)
}
