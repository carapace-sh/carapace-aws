package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var mturk_createHitCmd = &cobra.Command{
	Use:   "create-hit",
	Short: "The `CreateHIT` operation creates a new Human Intelligence Task (HIT).",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(mturk_createHitCmd).Standalone()

	mturk_createHitCmd.Flags().String("assignment-duration-in-seconds", "", "The amount of time, in seconds, that a Worker has to complete the HIT after accepting it.")
	mturk_createHitCmd.Flags().String("assignment-review-policy", "", "The Assignment-level Review Policy applies to the assignments under the HIT.")
	mturk_createHitCmd.Flags().String("auto-approval-delay-in-seconds", "", "The number of seconds after an assignment for the HIT has been submitted, after which the assignment is considered Approved automatically unless the Requester explicitly rejects it.")
	mturk_createHitCmd.Flags().String("description", "", "A general description of the HIT.")
	mturk_createHitCmd.Flags().String("hitlayout-id", "", "The HITLayoutId allows you to use a pre-existing HIT design with placeholder values and create an additional HIT by providing those values as HITLayoutParameters.")
	mturk_createHitCmd.Flags().String("hitlayout-parameters", "", "If the HITLayoutId is provided, any placeholder values must be filled in with values using the HITLayoutParameter structure.")
	mturk_createHitCmd.Flags().String("hitreview-policy", "", "The HIT-level Review Policy applies to the HIT.")
	mturk_createHitCmd.Flags().String("keywords", "", "One or more words or phrases that describe the HIT, separated by commas.")
	mturk_createHitCmd.Flags().String("lifetime-in-seconds", "", "An amount of time, in seconds, after which the HIT is no longer available for users to accept.")
	mturk_createHitCmd.Flags().String("max-assignments", "", "The number of times the HIT can be accepted and completed before the HIT becomes unavailable.")
	mturk_createHitCmd.Flags().String("qualification-requirements", "", "Conditions that a Worker's Qualifications must meet in order to accept the HIT.")
	mturk_createHitCmd.Flags().String("question", "", "The data the person completing the HIT uses to produce the results.")
	mturk_createHitCmd.Flags().String("requester-annotation", "", "An arbitrary data field.")
	mturk_createHitCmd.Flags().String("reward", "", "The amount of money the Requester will pay a Worker for successfully completing the HIT.")
	mturk_createHitCmd.Flags().String("title", "", "The title of the HIT.")
	mturk_createHitCmd.Flags().String("unique-request-token", "", "A unique identifier for this request which allows you to retry the call on error without creating duplicate HITs.")
	mturk_createHitCmd.MarkFlagRequired("assignment-duration-in-seconds")
	mturk_createHitCmd.MarkFlagRequired("description")
	mturk_createHitCmd.MarkFlagRequired("lifetime-in-seconds")
	mturk_createHitCmd.MarkFlagRequired("reward")
	mturk_createHitCmd.MarkFlagRequired("title")
	mturkCmd.AddCommand(mturk_createHitCmd)
}
