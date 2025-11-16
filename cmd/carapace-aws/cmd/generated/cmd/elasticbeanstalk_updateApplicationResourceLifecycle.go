package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_updateApplicationResourceLifecycleCmd = &cobra.Command{
	Use:   "update-application-resource-lifecycle",
	Short: "Modifies lifecycle settings for an application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_updateApplicationResourceLifecycleCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticbeanstalk_updateApplicationResourceLifecycleCmd).Standalone()

		elasticbeanstalk_updateApplicationResourceLifecycleCmd.Flags().String("application-name", "", "The name of the application.")
		elasticbeanstalk_updateApplicationResourceLifecycleCmd.Flags().String("resource-lifecycle-config", "", "The lifecycle configuration.")
		elasticbeanstalk_updateApplicationResourceLifecycleCmd.MarkFlagRequired("application-name")
		elasticbeanstalk_updateApplicationResourceLifecycleCmd.MarkFlagRequired("resource-lifecycle-config")
	})
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_updateApplicationResourceLifecycleCmd)
}
