package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_updateFleetCmd = &cobra.Command{
	Use:   "update-fleet",
	Short: "Updates a compute fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_updateFleetCmd).Standalone()

	codebuild_updateFleetCmd.Flags().String("arn", "", "The ARN of the compute fleet.")
	codebuild_updateFleetCmd.Flags().String("base-capacity", "", "The initial number of machines allocated to the compute ﬂeet, which deﬁnes the number of builds that can run in parallel.")
	codebuild_updateFleetCmd.Flags().String("compute-configuration", "", "The compute configuration of the compute fleet.")
	codebuild_updateFleetCmd.Flags().String("compute-type", "", "Information about the compute resources the compute fleet uses.")
	codebuild_updateFleetCmd.Flags().String("environment-type", "", "The environment type of the compute fleet.")
	codebuild_updateFleetCmd.Flags().String("fleet-service-role", "", "The service role associated with the compute fleet.")
	codebuild_updateFleetCmd.Flags().String("image-id", "", "The Amazon Machine Image (AMI) of the compute fleet.")
	codebuild_updateFleetCmd.Flags().String("overflow-behavior", "", "The compute fleet overflow behavior.")
	codebuild_updateFleetCmd.Flags().String("proxy-configuration", "", "The proxy configuration of the compute fleet.")
	codebuild_updateFleetCmd.Flags().String("scaling-configuration", "", "The scaling configuration of the compute fleet.")
	codebuild_updateFleetCmd.Flags().String("tags", "", "A list of tag key and value pairs associated with this compute fleet.")
	codebuild_updateFleetCmd.Flags().String("vpc-config", "", "")
	codebuild_updateFleetCmd.MarkFlagRequired("arn")
	codebuildCmd.AddCommand(codebuild_updateFleetCmd)
}
