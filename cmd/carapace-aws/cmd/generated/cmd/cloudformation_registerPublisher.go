package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloudformation_registerPublisherCmd = &cobra.Command{
	Use:   "register-publisher",
	Short: "Registers your account as a publisher of public extensions in the CloudFormation registry.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloudformation_registerPublisherCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(cloudformation_registerPublisherCmd).Standalone()

		cloudformation_registerPublisherCmd.Flags().String("accept-terms-and-conditions", "", "Whether you accept the [Terms and Conditions](https://cloudformation-registry-documents.s3.amazonaws.com/Terms_and_Conditions_for_AWS_CloudFormation_Registry_Publishers.pdf) for publishing extensions in the CloudFormation registry.")
		cloudformation_registerPublisherCmd.Flags().String("connection-arn", "", "If you are using a Bitbucket or GitHub account for identity verification, the Amazon Resource Name (ARN) for your connection to that account.")
	})
	cloudformationCmd.AddCommand(cloudformation_registerPublisherCmd)
}
