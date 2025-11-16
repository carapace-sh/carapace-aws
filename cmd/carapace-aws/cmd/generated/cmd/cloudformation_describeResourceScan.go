package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_describeResourceScanCmd = &cobra.Command{
	Use:   "describe-resource-scan",
	Short: "Describes details of a resource scan.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_describeResourceScanCmd).Standalone()

	cloudformation_describeResourceScanCmd.Flags().String("resource-scan-id", "", "The Amazon Resource Name (ARN) of the resource scan.")
	cloudformation_describeResourceScanCmd.MarkFlagRequired("resource-scan-id")
	cloudformationCmd.AddCommand(cloudformation_describeResourceScanCmd)
}
