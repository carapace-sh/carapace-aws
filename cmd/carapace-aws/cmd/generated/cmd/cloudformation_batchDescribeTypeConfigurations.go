package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_batchDescribeTypeConfigurationsCmd = &cobra.Command{
	Use:   "batch-describe-type-configurations",
	Short: "Returns configuration data for the specified CloudFormation extensions, from the CloudFormation registry in your current account and Region.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_batchDescribeTypeConfigurationsCmd).Standalone()

	cloudformation_batchDescribeTypeConfigurationsCmd.Flags().String("type-configuration-identifiers", "", "The list of identifiers for the desired extension configurations.")
	cloudformation_batchDescribeTypeConfigurationsCmd.MarkFlagRequired("type-configuration-identifiers")
	cloudformationCmd.AddCommand(cloudformation_batchDescribeTypeConfigurationsCmd)
}
