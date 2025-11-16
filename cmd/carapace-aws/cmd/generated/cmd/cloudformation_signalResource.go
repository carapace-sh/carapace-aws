package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_signalResourceCmd = &cobra.Command{
	Use:   "signal-resource",
	Short: "Sends a signal to the specified resource with a success or failure status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_signalResourceCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_signalResourceCmd).Standalone()

		cloudformation_signalResourceCmd.Flags().String("logical-resource-id", "", "The logical ID of the resource that you want to signal.")
		cloudformation_signalResourceCmd.Flags().String("stack-name", "", "The stack name or unique stack ID that includes the resource that you want to signal.")
		cloudformation_signalResourceCmd.Flags().String("status", "", "The status of the signal, which is either success or failure.")
		cloudformation_signalResourceCmd.Flags().String("unique-id", "", "A unique ID of the signal.")
		cloudformation_signalResourceCmd.MarkFlagRequired("logical-resource-id")
		cloudformation_signalResourceCmd.MarkFlagRequired("stack-name")
		cloudformation_signalResourceCmd.MarkFlagRequired("status")
		cloudformation_signalResourceCmd.MarkFlagRequired("unique-id")
	})
	cloudformationCmd.AddCommand(cloudformation_signalResourceCmd)
}
