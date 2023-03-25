package test

import (
	"fmt"
	"log"
	"time"

	"gitlab.utc.fr/wanhongz/emergency-simulator/agent"

	"gitlab.utc.fr/wanhongz/emergency-simulator/agent/rooms"
)

func Test() {
	h := agent.CreateHospital("localhost","8082")
	go h.Start()

	//log.Println(h.NurseCenter.GetNurseNumber())
	time.Sleep(5 * time.Second)
	//h.ReceptionCenter.ReduceQueue()
	//h.ReceptionCenter.ReduceQueue()
	//h.ReceptionCenter.ReduceQueue()
	//log.Println(h.ReceptionCenter.GetQueuesNumber())

	//log.Println(h.NurseCenter.GetNurseNumber())

	//time.Sleep(5*time.Second)

	//h.NurseCenter.Reduce_patient()
	//log.Println(h.NurseCenter.GetNurseNumber())
	h.EmergencyRoomCenter.AddRoom(3)
	h.EmergencyRoomCenter.AddRoom(3)
	h.EmergencyRoomCenter.AddRoom(3)

	fmt.Scanln()
}

func p(h *agent.Hospital) {
	time.Sleep(3 * time.Second)

	//h.AcceptNewPatient(1, true, "111", 10)
	//h.AcceptNewPatient(1, true, "111", 10)
	//h.AcceptNewPatient(1, true, "111", 10)
	//h.AcceptNewPatient(1, true, "111", 10)

	time.Sleep(50 * time.Second)
}

func Test2() {
	r := rooms.GetEmergencyRoomManagerInstance(10)
	go r.Run()
	n(r)
	fmt.Scanln()
}

func n(rr *rooms.EmergencyRoomManager) {
	c_request := rr.MsgRequest
	c_reponse := rr.MsgReponse
	c_request <- 3
	pp := <-c_reponse
	c_request <- 3
	pp2 := <-c_reponse

	log.Println(pp.ID)
	log.Println(pp2.ID)
}
