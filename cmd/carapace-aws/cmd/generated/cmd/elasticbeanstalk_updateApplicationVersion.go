package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_updateApplicationVersionCmd = &cobra.Command{
	Use:   "update-application-version",
	Short: "Updates the specified application version to have the specified properties.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_updateApplicationVersionCmd).Standalone()

	elasticbeanstalk_updateApplicationVersionCmd.Flags().String("application-name", "", "The name of the application associated with this version.")
	elasticbeanstalk_updateApplicationVersionCmd.Flags().String("description", "", "A new description for this version.")
	elasticbeanstalk_updateApplicationVersionCmd.Flags().String("version-label", "", "The name of the version to update.")
	elasticbeanstalk_updateApplicationVersionCmd.MarkFlagRequired("application-name")
	elasticbeanstalk_updateApplicationVersionCmd.MarkFlagRequired("version-label")
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_updateApplicationVersionCmd)
}
