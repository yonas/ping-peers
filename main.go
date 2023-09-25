package main

import (
  "broadcast/cmd/root"
  "broadcast/cmd/agent"
)

func main() {
  root.RootCmd().AddCommand(agent.AgentCmd())
  root.Execute()
}
