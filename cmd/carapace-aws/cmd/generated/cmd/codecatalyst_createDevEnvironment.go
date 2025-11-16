package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var codecatalyst_createDevEnvironmentCmd = &cobra.Command{
	Use:   "create-dev-environment",
	Short: "Creates a Dev Environment in Amazon CodeCatalyst, a cloud-based development environment that you can use to quickly work on the code stored in the source repositories of your project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(codecatalyst_createDevEnvironmentCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(codecatalyst_createDevEnvironmentCmd).Standalone()

		codecatalyst_createDevEnvironmentCmd.Flags().String("alias", "", "The user-defined alias for a Dev Environment.")
		codecatalyst_createDevEnvironmentCmd.Flags().String("client-token", "", "A user-specified idempotency token.")
		codecatalyst_createDevEnvironmentCmd.Flags().String("ides", "", "Information about the integrated development environment (IDE) configured for a Dev Environment.")
		codecatalyst_createDevEnvironmentCmd.Flags().String("inactivity-timeout-minutes", "", "The amount of time the Dev Environment will run without any activity detected before stopping, in minutes.")
		codecatalyst_createDevEnvironmentCmd.Flags().String("instance-type", "", "The Amazon EC2 instace type to use for the Dev Environment.")
		codecatalyst_createDevEnvironmentCmd.Flags().String("persistent-storage", "", "Information about the amount of storage allocated to the Dev Environment.")
		codecatalyst_createDevEnvironmentCmd.Flags().String("project-name", "", "The name of the project in the space.")
		codecatalyst_createDevEnvironmentCmd.Flags().String("repositories", "", "The source repository that contains the branch to clone into the Dev Environment.")
		codecatalyst_createDevEnvironmentCmd.Flags().String("space-name", "", "The name of the space.")
		codecatalyst_createDevEnvironmentCmd.Flags().String("vpc-connection-name", "", "The name of the connection that will be used to connect to Amazon VPC, if any.")
		codecatalyst_createDevEnvironmentCmd.MarkFlagRequired("instance-type")
		codecatalyst_createDevEnvironmentCmd.MarkFlagRequired("persistent-storage")
		codecatalyst_createDevEnvironmentCmd.MarkFlagRequired("project-name")
		codecatalyst_createDevEnvironmentCmd.MarkFlagRequired("space-name")
	})
	codecatalystCmd.AddCommand(codecatalyst_createDevEnvironmentCmd)
}
