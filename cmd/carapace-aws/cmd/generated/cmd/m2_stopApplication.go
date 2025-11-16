package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var m2_stopApplicationCmd = &cobra.Command{
	Use:   "stop-application",
	Short: "Stops a running application.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(m2_stopApplicationCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(m2_stopApplicationCmd).Standalone()

		m2_stopApplicationCmd.Flags().String("application-id", "", "The unique identifier of the application you want to stop.")
		m2_stopApplicationCmd.Flags().Bool("force-stop", false, "Stopping an application process can take a long time.")
		m2_stopApplicationCmd.Flags().Bool("no-force-stop", false, "Stopping an application process can take a long time.")
		m2_stopApplicationCmd.MarkFlagRequired("application-id")
		m2_stopApplicationCmd.Flag("no-force-stop").Hidden = true
	})
	m2Cmd.AddCommand(m2_stopApplicationCmd)
}
