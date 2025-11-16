package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_updateTestGridProjectCmd = &cobra.Command{
	Use:   "update-test-grid-project",
	Short: "Change details of a project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_updateTestGridProjectCmd).Standalone()

	devicefarm_updateTestGridProjectCmd.Flags().String("description", "", "Human-readable description for the project.")
	devicefarm_updateTestGridProjectCmd.Flags().String("name", "", "Human-readable name for the project.")
	devicefarm_updateTestGridProjectCmd.Flags().String("project-arn", "", "ARN of the project to update.")
	devicefarm_updateTestGridProjectCmd.Flags().String("vpc-config", "", "The VPC security groups and subnets that are attached to a project.")
	devicefarm_updateTestGridProjectCmd.MarkFlagRequired("project-arn")
	devicefarmCmd.AddCommand(devicefarm_updateTestGridProjectCmd)
}
