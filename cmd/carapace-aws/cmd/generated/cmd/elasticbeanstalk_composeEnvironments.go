package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var elasticbeanstalk_composeEnvironmentsCmd = &cobra.Command{
	Use:   "compose-environments",
	Short: "Create or update a group of environments that each run a separate component of a single application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(elasticbeanstalk_composeEnvironmentsCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(elasticbeanstalk_composeEnvironmentsCmd).Standalone()

		elasticbeanstalk_composeEnvironmentsCmd.Flags().String("application-name", "", "The name of the application to which the specified source bundles belong.")
		elasticbeanstalk_composeEnvironmentsCmd.Flags().String("group-name", "", "The name of the group to which the target environments belong.")
		elasticbeanstalk_composeEnvironmentsCmd.Flags().String("version-labels", "", "A list of version labels, specifying one or more application source bundles that belong to the target application.")
	})
	elasticbeanstalkCmd.AddCommand(elasticbeanstalk_composeEnvironmentsCmd)
}
