package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_createProjectCmd = &cobra.Command{
	Use:   "create-project",
	Short: "Creates a project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_createProjectCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_createProjectCmd).Standalone()

		devicefarm_createProjectCmd.Flags().String("default-job-timeout-minutes", "", "Sets the execution timeout value (in minutes) for a project.")
		devicefarm_createProjectCmd.Flags().String("name", "", "The project's name.")
		devicefarm_createProjectCmd.Flags().String("vpc-config", "", "The VPC security groups and subnets that are attached to a project.")
		devicefarm_createProjectCmd.MarkFlagRequired("name")
	})
	devicefarmCmd.AddCommand(devicefarm_createProjectCmd)
}
