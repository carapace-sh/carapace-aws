package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_testTypeCmd = &cobra.Command{
	Use:   "test-type",
	Short: "Tests a registered extension to make sure it meets all necessary requirements for being published in the CloudFormation registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_testTypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_testTypeCmd).Standalone()

		cloudformation_testTypeCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the extension.")
		cloudformation_testTypeCmd.Flags().String("log-delivery-bucket", "", "The S3 bucket to which CloudFormation delivers the contract test execution logs.")
		cloudformation_testTypeCmd.Flags().String("type", "", "The type of the extension to test.")
		cloudformation_testTypeCmd.Flags().String("type-name", "", "The name of the extension to test.")
		cloudformation_testTypeCmd.Flags().String("version-id", "", "The version of the extension to test.")
	})
	cloudformationCmd.AddCommand(cloudformation_testTypeCmd)
}
