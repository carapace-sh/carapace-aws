package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fms_getResourceSetCmd = &cobra.Command{
	Use:   "get-resource-set",
	Short: "Gets information about a specific resource set.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fms_getResourceSetCmd).Standalone()

	fms_getResourceSetCmd.Flags().String("identifier", "", "A unique identifier for the resource set, used in a request to refer to the resource set.")
	fms_getResourceSetCmd.MarkFlagRequired("identifier")
	fmsCmd.AddCommand(fms_getResourceSetCmd)
}
