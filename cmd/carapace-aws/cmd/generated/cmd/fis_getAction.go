package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fis_getActionCmd = &cobra.Command{
	Use:   "get-action",
	Short: "Gets information about the specified FIS action.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fis_getActionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fis_getActionCmd).Standalone()

		fis_getActionCmd.Flags().String("id", "", "The ID of the action.")
		fis_getActionCmd.MarkFlagRequired("id")
	})
	fisCmd.AddCommand(fis_getActionCmd)
}
