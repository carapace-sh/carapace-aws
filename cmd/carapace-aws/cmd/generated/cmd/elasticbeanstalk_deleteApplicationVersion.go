package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_deleteApplicationVersionCmd = &cobra.Command{
	Use:   "delete-application-version",
	Short: "Deletes the specified version from the specified application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_deleteApplicationVersionCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticbeanstalk_deleteApplicationVersionCmd).Standalone()

		elasticbeanstalk_deleteApplicationVersionCmd.Flags().String("application-name", "", "The name of the application to which the version belongs.")
		elasticbeanstalk_deleteApplicationVersionCmd.Flags().String("delete-source-bundle", "", "Set to `true` to delete the source bundle from your storage bucket.")
		elasticbeanstalk_deleteApplicationVersionCmd.Flags().String("version-label", "", "The label of the version to delete.")
		elasticbeanstalk_deleteApplicationVersionCmd.MarkFlagRequired("application-name")
		elasticbeanstalk_deleteApplicationVersionCmd.MarkFlagRequired("version-label")
	})
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_deleteApplicationVersionCmd)
}
