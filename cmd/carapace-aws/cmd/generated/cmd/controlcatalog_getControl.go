package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var controlcatalog_getControlCmd = &cobra.Command{
	Use:   "get-control",
	Short: "Returns details about a specific control, most notably a list of Amazon Web Services Regions where this control is supported.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(controlcatalog_getControlCmd).Standalone()

	controlcatalog_getControlCmd.Flags().String("control-arn", "", "The Amazon Resource Name (ARN) of the control.")
	controlcatalog_getControlCmd.MarkFlagRequired("control-arn")
	controlcatalogCmd.AddCommand(controlcatalog_getControlCmd)
}
