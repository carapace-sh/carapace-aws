package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_deleteApplicationCmd = &cobra.Command{
	Use:   "delete-application",
	Short: "Deletes the specified application along with all associated versions and configurations.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_deleteApplicationCmd).Standalone()

	elasticbeanstalk_deleteApplicationCmd.Flags().String("application-name", "", "The name of the application to delete.")
	elasticbeanstalk_deleteApplicationCmd.Flags().String("terminate-env-by-force", "", "When set to true, running environments will be terminated before deleting the application.")
	elasticbeanstalk_deleteApplicationCmd.MarkFlagRequired("application-name")
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_deleteApplicationCmd)
}
