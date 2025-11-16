package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var devicefarm_createTestGridUrlCmd = &cobra.Command{
	Use:   "create-test-grid-url",
	Short: "Creates a signed, short-term URL that can be passed to a Selenium `RemoteWebDriver` constructor.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(devicefarm_createTestGridUrlCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(devicefarm_createTestGridUrlCmd).Standalone()

		devicefarm_createTestGridUrlCmd.Flags().String("expires-in-seconds", "", "Lifetime, in seconds, of the URL.")
		devicefarm_createTestGridUrlCmd.Flags().String("project-arn", "", "ARN (from [CreateTestGridProject]() or [ListTestGridProjects]()) to associate with the short-term URL.")
		devicefarm_createTestGridUrlCmd.MarkFlagRequired("expires-in-seconds")
		devicefarm_createTestGridUrlCmd.MarkFlagRequired("project-arn")
	})
	devicefarmCmd.AddCommand(devicefarm_createTestGridUrlCmd)
}
