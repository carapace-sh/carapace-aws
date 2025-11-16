package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_createHittypeCmd = &cobra.Command{
	Use:   "create-hittype",
	Short: "The `CreateHITType` operation creates a new HIT type.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_createHittypeCmd).PreRun(func(cmd *cobra.Command, args []string) {
		carapace.Gen(mturk_createHittypeCmd).Standalone()

		mturk_createHittypeCmd.Flags().String("assignment-duration-in-seconds", "", "The amount of time, in seconds, that a Worker has to complete the HIT after accepting it.")
		mturk_createHittypeCmd.Flags().String("auto-approval-delay-in-seconds", "", "The number of seconds after an assignment for the HIT has been submitted, after which the assignment is considered Approved automatically unless the Requester explicitly rejects it.")
		mturk_createHittypeCmd.Flags().String("description", "", "A general description of the HIT.")
		mturk_createHittypeCmd.Flags().String("keywords", "", "One or more words or phrases that describe the HIT, separated by commas.")
		mturk_createHittypeCmd.Flags().String("qualification-requirements", "", "Conditions that a Worker's Qualifications must meet in order to accept the HIT.")
		mturk_createHittypeCmd.Flags().String("reward", "", "The amount of money the Requester will pay a Worker for successfully completing the HIT.")
		mturk_createHittypeCmd.Flags().String("title", "", "The title of the HIT.")
		mturk_createHittypeCmd.MarkFlagRequired("assignment-duration-in-seconds")
		mturk_createHittypeCmd.MarkFlagRequired("description")
		mturk_createHittypeCmd.MarkFlagRequired("reward")
		mturk_createHittypeCmd.MarkFlagRequired("title")
	})
	mturkCmd.AddCommand(mturk_createHittypeCmd)
}
