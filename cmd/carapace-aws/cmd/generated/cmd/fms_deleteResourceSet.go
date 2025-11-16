package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_deleteResourceSetCmd = &cobra.Command{
	Use:   "delete-resource-set",
	Short: "Deletes the specified [ResourceSet]().",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_deleteResourceSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fms_deleteResourceSetCmd).Standalone()

		fms_deleteResourceSetCmd.Flags().String("identifier", "", "A unique identifier for the resource set, used in a request to refer to the resource set.")
		fms_deleteResourceSetCmd.MarkFlagRequired("identifier")
	})
	fmsCmd.AddCommand(fms_deleteResourceSetCmd)
}
