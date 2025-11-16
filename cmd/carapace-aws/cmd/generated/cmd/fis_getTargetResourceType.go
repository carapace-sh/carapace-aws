package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var fis_getTargetResourceTypeCmd = &cobra.Command{
	Use:   "get-target-resource-type",
	Short: "Gets information about the specified resource type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(fis_getTargetResourceTypeCmd).Standalone()

	fis_getTargetResourceTypeCmd.Flags().String("resource-type", "", "The resource type.")
	fis_getTargetResourceTypeCmd.MarkFlagRequired("resource-type")
	fisCmd.AddCommand(fis_getTargetResourceTypeCmd)
}
