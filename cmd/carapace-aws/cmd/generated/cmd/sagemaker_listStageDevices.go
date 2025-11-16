package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var sagemaker_listStageDevicesCmd = &cobra.Command{
	Use:   "list-stage-devices",
	Short: "Lists devices allocated to the stage, containing detailed device information and deployment status.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(sagemaker_listStageDevicesCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(sagemaker_listStageDevicesCmd).Standalone()

		sagemaker_listStageDevicesCmd.Flags().String("edge-deployment-plan-name", "", "The name of the edge deployment plan.")
		sagemaker_listStageDevicesCmd.Flags().Bool("exclude-devices-deployed-in-other-stage", false, "Toggle for excluding devices deployed in other stages.")
		sagemaker_listStageDevicesCmd.Flags().String("max-results", "", "The maximum number of requests to select.")
		sagemaker_listStageDevicesCmd.Flags().String("next-token", "", "The response from the last list when returning a list large enough to neeed tokening.")
		sagemaker_listStageDevicesCmd.Flags().Bool("no-exclude-devices-deployed-in-other-stage", false, "Toggle for excluding devices deployed in other stages.")
		sagemaker_listStageDevicesCmd.Flags().String("stage-name", "", "The name of the stage in the deployment.")
		sagemaker_listStageDevicesCmd.MarkFlagRequired("edge-deployment-plan-name")
		sagemaker_listStageDevicesCmd.Flag("no-exclude-devices-deployed-in-other-stage").Hidden = true
		sagemaker_listStageDevicesCmd.MarkFlagRequired("stage-name")
	})
	sagemakerCmd.AddCommand(sagemaker_listStageDevicesCmd)
}
