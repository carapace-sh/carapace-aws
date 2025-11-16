package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var comprehend_listFlywheelsCmd = &cobra.Command{
	Use:   "list-flywheels",
	Short: "Gets a list of the flywheels that you have created.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(comprehend_listFlywheelsCmd).Standalone()

	comprehend_listFlywheelsCmd.Flags().String("filter", "", "Filters the flywheels that are returned.")
	comprehend_listFlywheelsCmd.Flags().String("max-results", "", "Maximum number of results to return in a response.")
	comprehend_listFlywheelsCmd.Flags().String("next-token", "", "Identifies the next page of results to return.")
	comprehendCmd.AddCommand(comprehend_listFlywheelsCmd)
}
