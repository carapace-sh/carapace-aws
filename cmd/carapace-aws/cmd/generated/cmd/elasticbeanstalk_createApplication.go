package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_createApplicationCmd = &cobra.Command{
	Use:   "create-application",
	Short: "Creates an application that has one configuration template named `default` and no application versions.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_createApplicationCmd).Standalone()

	elasticbeanstalk_createApplicationCmd.Flags().String("application-name", "", "The name of the application.")
	elasticbeanstalk_createApplicationCmd.Flags().String("description", "", "Your description of the application.")
	elasticbeanstalk_createApplicationCmd.Flags().String("resource-lifecycle-config", "", "Specifies an application resource lifecycle configuration to prevent your application from accumulating too many versions.")
	elasticbeanstalk_createApplicationCmd.Flags().String("tags", "", "Specifies the tags applied to the application.")
	elasticbeanstalk_createApplicationCmd.MarkFlagRequired("application-name")
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_createApplicationCmd)
}
