package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mediastoreData_describeObjectCmd = &cobra.Command{
	Use:   "describe-object",
	Short: "Gets the headers for an object at the specified path.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mediastoreData_describeObjectCmd).Standalone()

	mediastoreData_describeObjectCmd.Flags().String("path", "", "The path (including the file name) where the object is stored in the container.")
	mediastoreData_describeObjectCmd.MarkFlagRequired("path")
	mediastoreDataCmd.AddCommand(mediastoreData_describeObjectCmd)
}
