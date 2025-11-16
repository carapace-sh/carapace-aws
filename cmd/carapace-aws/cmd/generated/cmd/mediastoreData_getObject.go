package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediastoreData_getObjectCmd = &cobra.Command{
	Use:   "get-object",
	Short: "Downloads the object at the specified path.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediastoreData_getObjectCmd).Standalone()

	mediastoreData_getObjectCmd.Flags().String("path", "", "The path (including the file name) where the object is stored in the container.")
	mediastoreData_getObjectCmd.Flags().String("range", "", "The range bytes of an object to retrieve.")
	mediastoreData_getObjectCmd.MarkFlagRequired("path")
	mediastoreDataCmd.AddCommand(mediastoreData_getObjectCmd)
}
