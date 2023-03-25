package main

import "gitlab.utc.fr/wanhongz/emergency-simulator/agent"

func main() {
	var h *agent.Hospital = agent.CreateHospital("127.0.0.1","8082")
	h.Start()
}
