package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var connect_resumeContactCmd = &cobra.Command{
	Use:   "resume-contact",
	Short: "Allows resuming a task contact in a paused state.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(connect_resumeContactCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(connect_resumeContactCmd).Standalone()

		connect_resumeContactCmd.Flags().String("contact-flow-id", "", "The identifier of the flow.")
		connect_resumeContactCmd.Flags().String("contact-id", "", "The identifier of the contact.")
		connect_resumeContactCmd.Flags().String("instance-id", "", "The identifier of the Amazon Connect instance.")
		connect_resumeContactCmd.MarkFlagRequired("contact-id")
		connect_resumeContactCmd.MarkFlagRequired("instance-id")
	})
	connectCmd.AddCommand(connect_resumeContactCmd)
}
