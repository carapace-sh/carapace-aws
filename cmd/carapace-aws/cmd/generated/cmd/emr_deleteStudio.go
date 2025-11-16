package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var emr_deleteStudioCmd = &cobra.Command{
	Use:   "delete-studio",
	Short: "Removes an Amazon EMR Studio from the Studio metadata store.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(emr_deleteStudioCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(emr_deleteStudioCmd).Standalone()

		emr_deleteStudioCmd.Flags().String("studio-id", "", "The ID of the Amazon EMR Studio.")
		emr_deleteStudioCmd.MarkFlagRequired("studio-id")
	})
	emrCmd.AddCommand(emr_deleteStudioCmd)
}
