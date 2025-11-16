package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_createTestGridProjectCmd = &cobra.Command{
	Use:   "create-test-grid-project",
	Short: "Creates a Selenium testing project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_createTestGridProjectCmd).Standalone()

	devicefarm_createTestGridProjectCmd.Flags().String("description", "", "Human-readable description of the project.")
	devicefarm_createTestGridProjectCmd.Flags().String("name", "", "Human-readable name of the Selenium testing project.")
	devicefarm_createTestGridProjectCmd.Flags().String("vpc-config", "", "The VPC security groups and subnets that are attached to a project.")
	devicefarm_createTestGridProjectCmd.MarkFlagRequired("name")
	devicefarmCmd.AddCommand(devicefarm_createTestGridProjectCmd)
}
