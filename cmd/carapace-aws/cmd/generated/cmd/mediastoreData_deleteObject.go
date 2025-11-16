package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediastoreData_deleteObjectCmd = &cobra.Command{
	Use:   "delete-object",
	Short: "Deletes an object at the specified path.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediastoreData_deleteObjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mediastoreData_deleteObjectCmd).Standalone()

		mediastoreData_deleteObjectCmd.Flags().String("path", "", "The path (including the file name) where the object is stored in the container.")
		mediastoreData_deleteObjectCmd.MarkFlagRequired("path")
	})
	mediastoreDataCmd.AddCommand(mediastoreData_deleteObjectCmd)
}
