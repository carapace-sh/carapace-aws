package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_updateApplicationCmd = &cobra.Command{
	Use:   "update-application",
	Short: "Updates the specified application to have the specified properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_updateApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticbeanstalk_updateApplicationCmd).Standalone()

		elasticbeanstalk_updateApplicationCmd.Flags().String("application-name", "", "The name of the application to update.")
		elasticbeanstalk_updateApplicationCmd.Flags().String("description", "", "A new description for the application.")
		elasticbeanstalk_updateApplicationCmd.MarkFlagRequired("application-name")
	})
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_updateApplicationCmd)
}
