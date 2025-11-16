package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_putResourceSetCmd = &cobra.Command{
	Use:   "put-resource-set",
	Short: "Creates the resource set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_putResourceSetCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(fms_putResourceSetCmd).Standalone()

		fms_putResourceSetCmd.Flags().String("resource-set", "", "Details about the resource set to be created or updated.&gt;")
		fms_putResourceSetCmd.Flags().String("tag-list", "", "Retrieves the tags associated with the specified resource set.")
		fms_putResourceSetCmd.MarkFlagRequired("resource-set")
	})
	fmsCmd.AddCommand(fms_putResourceSetCmd)
}
