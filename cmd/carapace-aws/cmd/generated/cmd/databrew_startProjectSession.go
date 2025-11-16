package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var databrew_startProjectSessionCmd = &cobra.Command{
	Use:   "start-project-session",
	Short: "Creates an interactive session, enabling you to manipulate data in a DataBrew project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(databrew_startProjectSessionCmd).Standalone()

	databrew_startProjectSessionCmd.Flags().String("assume-control", "", "A value that, if true, enables you to take control of a session, even if a different client is currently accessing the project.")
	databrew_startProjectSessionCmd.Flags().String("name", "", "The name of the project to act upon.")
	databrew_startProjectSessionCmd.MarkFlagRequired("name")
	databrewCmd.AddCommand(databrew_startProjectSessionCmd)
}
