package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var globalaccelerator_updateCustomRoutingAcceleratorAttributesCmd = &cobra.Command{
	Use:   "update-custom-routing-accelerator-attributes",
	Short: "Update the attributes for a custom routing accelerator.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(globalaccelerator_updateCustomRoutingAcceleratorAttributesCmd).Standalone()

	globalaccelerator_updateCustomRoutingAcceleratorAttributesCmd.Flags().String("accelerator-arn", "", "The Amazon Resource Name (ARN) of the custom routing accelerator to update attributes for.")
	globalaccelerator_updateCustomRoutingAcceleratorAttributesCmd.Flags().String("flow-logs-enabled", "", "Update whether flow logs are enabled.")
	globalaccelerator_updateCustomRoutingAcceleratorAttributesCmd.Flags().String("flow-logs-s3-bucket", "", "The name of the Amazon S3 bucket for the flow logs.")
	globalaccelerator_updateCustomRoutingAcceleratorAttributesCmd.Flags().String("flow-logs-s3-prefix", "", "Update the prefix for the location in the Amazon S3 bucket for the flow logs.")
	globalaccelerator_updateCustomRoutingAcceleratorAttributesCmd.MarkFlagRequired("accelerator-arn")
	globalacceleratorCmd.AddCommand(globalaccelerator_updateCustomRoutingAcceleratorAttributesCmd)
}
