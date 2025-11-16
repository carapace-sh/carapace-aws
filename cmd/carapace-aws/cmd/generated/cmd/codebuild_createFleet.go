package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codebuild_createFleetCmd = &cobra.Command{
	Use:   "create-fleet",
	Short: "Creates a compute fleet.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codebuild_createFleetCmd).Standalone()

	codebuild_createFleetCmd.Flags().String("base-capacity", "", "The initial number of machines allocated to the ﬂeet, which deﬁnes the number of builds that can run in parallel.")
	codebuild_createFleetCmd.Flags().String("compute-configuration", "", "The compute configuration of the compute fleet.")
	codebuild_createFleetCmd.Flags().String("compute-type", "", "Information about the compute resources the compute fleet uses.")
	codebuild_createFleetCmd.Flags().String("environment-type", "", "The environment type of the compute fleet.")
	codebuild_createFleetCmd.Flags().String("fleet-service-role", "", "The service role associated with the compute fleet.")
	codebuild_createFleetCmd.Flags().String("image-id", "", "The Amazon Machine Image (AMI) of the compute fleet.")
	codebuild_createFleetCmd.Flags().String("name", "", "The name of the compute fleet.")
	codebuild_createFleetCmd.Flags().String("overflow-behavior", "", "The compute fleet overflow behavior.")
	codebuild_createFleetCmd.Flags().String("proxy-configuration", "", "The proxy configuration of the compute fleet.")
	codebuild_createFleetCmd.Flags().String("scaling-configuration", "", "The scaling configuration of the compute fleet.")
	codebuild_createFleetCmd.Flags().String("tags", "", "A list of tag key and value pairs associated with this compute fleet.")
	codebuild_createFleetCmd.Flags().String("vpc-config", "", "")
	codebuild_createFleetCmd.MarkFlagRequired("base-capacity")
	codebuild_createFleetCmd.MarkFlagRequired("compute-type")
	codebuild_createFleetCmd.MarkFlagRequired("environment-type")
	codebuild_createFleetCmd.MarkFlagRequired("name")
	codebuildCmd.AddCommand(codebuild_createFleetCmd)
}
