package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_updateAcceleratorAttributesCmd = &cobra.Command{
	Use:   "update-accelerator-attributes",
	Short: "Update the attributes for an accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_updateAcceleratorAttributesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(globalaccelerator_updateAcceleratorAttributesCmd).Standalone()

		globalaccelerator_updateAcceleratorAttributesCmd.Flags().String("accelerator-arn", "", "The Amazon Resource Name (ARN) of the accelerator that you want to update.")
		globalaccelerator_updateAcceleratorAttributesCmd.Flags().String("flow-logs-enabled", "", "Update whether flow logs are enabled.")
		globalaccelerator_updateAcceleratorAttributesCmd.Flags().String("flow-logs-s3-bucket", "", "The name of the Amazon S3 bucket for the flow logs.")
		globalaccelerator_updateAcceleratorAttributesCmd.Flags().String("flow-logs-s3-prefix", "", "Update the prefix for the location in the Amazon S3 bucket for the flow logs.")
		globalaccelerator_updateAcceleratorAttributesCmd.MarkFlagRequired("accelerator-arn")
	})
	globalacceleratorCmd.AddCommand(globalaccelerator_updateAcceleratorAttributesCmd)
}
