package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_updateProjectCmd = &cobra.Command{
	Use:   "update-project",
	Short: "Modifies the specified project name, given the project ARN and a new name.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_updateProjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_updateProjectCmd).Standalone()

		devicefarm_updateProjectCmd.Flags().String("arn", "", "The Amazon Resource Name (ARN) of the project whose name to update.")
		devicefarm_updateProjectCmd.Flags().String("default-job-timeout-minutes", "", "The number of minutes a test run in the project executes before it times out.")
		devicefarm_updateProjectCmd.Flags().String("name", "", "A string that represents the new name of the project that you are updating.")
		devicefarm_updateProjectCmd.Flags().String("vpc-config", "", "The VPC security groups and subnets that are attached to a project.")
		devicefarm_updateProjectCmd.MarkFlagRequired("arn")
	})
	devicefarmCmd.AddCommand(devicefarm_updateProjectCmd)
}
