package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var lightsail_createCloudFormationStackCmd = &cobra.Command{
	Use:   "create-cloud-formation-stack",
	Short: "Creates an AWS CloudFormation stack, which creates a new Amazon EC2 instance from an exported Amazon Lightsail snapshot.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(lightsail_createCloudFormationStackCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(lightsail_createCloudFormationStackCmd).Standalone()

		lightsail_createCloudFormationStackCmd.Flags().String("instances", "", "An array of parameters that will be used to create the new Amazon EC2 instance.")
		lightsail_createCloudFormationStackCmd.MarkFlagRequired("instances")
	})
	lightsailCmd.AddCommand(lightsail_createCloudFormationStackCmd)
}
