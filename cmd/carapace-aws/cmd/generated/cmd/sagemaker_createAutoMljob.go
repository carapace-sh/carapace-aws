package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_createAutoMljobCmd = &cobra.Command{
	Use:   "create-auto-mljob",
	Short: "Creates an Autopilot job also referred to as Autopilot experiment or AutoML job.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_createAutoMljobCmd).Standalone()

	sagemaker_createAutoMljobCmd.Flags().String("auto-mljob-config", "", "A collection of settings used to configure an AutoML job.")
	sagemaker_createAutoMljobCmd.Flags().String("auto-mljob-name", "", "Identifies an Autopilot job.")
	sagemaker_createAutoMljobCmd.Flags().String("auto-mljob-objective", "", "Specifies a metric to minimize or maximize as the objective of a job.")
	sagemaker_createAutoMljobCmd.Flags().String("generate-candidate-definitions-only", "", "Generates possible candidates without training the models.")
	sagemaker_createAutoMljobCmd.Flags().String("input-data-config", "", "An array of channel objects that describes the input data and its location.")
	sagemaker_createAutoMljobCmd.Flags().String("model-deploy-config", "", "Specifies how to generate the endpoint name for an automatic one-click Autopilot model deployment.")
	sagemaker_createAutoMljobCmd.Flags().String("output-data-config", "", "Provides information about encryption and the Amazon S3 output path needed to store artifacts from an AutoML job.")
	sagemaker_createAutoMljobCmd.Flags().String("problem-type", "", "Defines the type of supervised learning problem available for the candidates.")
	sagemaker_createAutoMljobCmd.Flags().String("role-arn", "", "The ARN of the role that is used to access the data.")
	sagemaker_createAutoMljobCmd.Flags().String("tags", "", "An array of key-value pairs.")
	sagemaker_createAutoMljobCmd.MarkFlagRequired("auto-mljob-name")
	sagemaker_createAutoMljobCmd.MarkFlagRequired("input-data-config")
	sagemaker_createAutoMljobCmd.MarkFlagRequired("output-data-config")
	sagemaker_createAutoMljobCmd.MarkFlagRequired("role-arn")
	sagemakerCmd.AddCommand(sagemaker_createAutoMljobCmd)
}
