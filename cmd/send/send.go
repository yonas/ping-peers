package send

import (
  "broadcast/pkg/send"

  "github.com/spf13/cobra"
)

var sendTags string

var sendCmd = &cobra.Command{
  Use:   "send",
  Aliases: []string{"s"},
  Short:  "Send a string",
  //Args:  cobra.ExactArgs(1),
  Run: func(cmd *cobra.Command, args []string) {
    send.broadcast(sendTags)
  },
}

func init() {
  sendCmd.Flags().StringVarP(&sendTags, "tags", "t", "", "Tags")
  rootCmd.AddCommand(sendCmd)
}
