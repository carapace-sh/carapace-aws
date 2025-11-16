package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var finspaceData_getWorkingLocationCmd = &cobra.Command{
	Use:   "get-working-location",
	Short: "A temporary Amazon S3 location, where you can copy your files from a source location to stage or use as a scratch space in FinSpace notebook.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(finspaceData_getWorkingLocationCmd).Standalone()

	finspaceData_getWorkingLocationCmd.Flags().String("location-type", "", "Specify the type of the working location.")
	finspaceDataCmd.AddCommand(finspaceData_getWorkingLocationCmd)
}
