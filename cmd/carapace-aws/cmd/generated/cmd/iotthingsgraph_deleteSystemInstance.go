package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var iotthingsgraph_deleteSystemInstanceCmd = &cobra.Command{
	Use:   "delete-system-instance",
	Short: "Deletes a system instance.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(iotthingsgraph_deleteSystemInstanceCmd).Standalone()

	iotthingsgraph_deleteSystemInstanceCmd.Flags().String("id", "", "The ID of the system instance to be deleted.")
	iotthingsgraphCmd.AddCommand(iotthingsgraph_deleteSystemInstanceCmd)
}
