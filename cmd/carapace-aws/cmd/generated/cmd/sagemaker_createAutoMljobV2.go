package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createAutoMljobV2Cmd = &cobra.Command{
	Use:   "create-auto-mljob-v2",
	Short: "Creates an Autopilot job also referred to as Autopilot experiment or AutoML job V2.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createAutoMljobV2Cmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_createAutoMljobV2Cmd).Standalone()

		sagemaker_createAutoMljobV2Cmd.Flags().String("auto-mlcompute-config", "", "Specifies the compute configuration for the AutoML job V2.")
		sagemaker_createAutoMljobV2Cmd.Flags().String("auto-mljob-input-data-config", "", "An array of channel objects describing the input data and their location.")
		sagemaker_createAutoMljobV2Cmd.Flags().String("auto-mljob-name", "", "Identifies an Autopilot job.")
		sagemaker_createAutoMljobV2Cmd.Flags().String("auto-mljob-objective", "", "Specifies a metric to minimize or maximize as the objective of a job.")
		sagemaker_createAutoMljobV2Cmd.Flags().String("auto-mlproblem-type-config", "", "Defines the configuration settings of one of the supported problem types.")
		sagemaker_createAutoMljobV2Cmd.Flags().String("data-split-config", "", "This structure specifies how to split the data into train and validation datasets.")
		sagemaker_createAutoMljobV2Cmd.Flags().String("model-deploy-config", "", "Specifies how to generate the endpoint name for an automatic one-click Autopilot model deployment.")
		sagemaker_createAutoMljobV2Cmd.Flags().String("output-data-config", "", "Provides information about encryption and the Amazon S3 output path needed to store artifacts from an AutoML job.")
		sagemaker_createAutoMljobV2Cmd.Flags().String("role-arn", "", "The ARN of the role that is used to access the data.")
		sagemaker_createAutoMljobV2Cmd.Flags().String("security-config", "", "The security configuration for traffic encryption or Amazon VPC settings.")
		sagemaker_createAutoMljobV2Cmd.Flags().String("tags", "", "An array of key-value pairs.")
		sagemaker_createAutoMljobV2Cmd.MarkFlagRequired("auto-mljob-input-data-config")
		sagemaker_createAutoMljobV2Cmd.MarkFlagRequired("auto-mljob-name")
		sagemaker_createAutoMljobV2Cmd.MarkFlagRequired("auto-mlproblem-type-config")
		sagemaker_createAutoMljobV2Cmd.MarkFlagRequired("output-data-config")
		sagemaker_createAutoMljobV2Cmd.MarkFlagRequired("role-arn")
	})
	sagemakerCmd.AddCommand(sagemaker_createAutoMljobV2Cmd)
}
