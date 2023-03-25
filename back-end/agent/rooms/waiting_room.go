package rooms

import (
	"log"
	"strconv"
	"sync"
	"time"

	"gitlab.utc.fr/wanhongz/emergency-simulator/agent/doctor"
	"gitlab.utc.fr/wanhongz/emergency-simulator/agent/patient"
)

type WaitingRoom struct {
	sync.Mutex
	MsgRequest           chan *patient.Patient // Le canal qui accepte la demande du patient, après avoir écouté la demande du patient, l'ajoute à la file d'attente
	QueuePatients        []*patient.Patient    // tous les patients en attente
	EmergencyRoomRequest chan int              // Canaux pour demander des ressources en salle d'urgence
	EmergencyRoomReponse chan *EmergencyRoom   // Canal de rétroaction pour les ressources des salles d'urgence
	DoctorsRequest       chan int              // Postuler pour la chaîne du médecin
	DoctorsResponse      chan *doctor.Doctor
	Alltime              map[string]float64
}

var (
	instance_wr *WaitingRoom
	once3       sync.Once
)

func GetWaitingRoomInstance(c1 chan int, c2 chan *EmergencyRoom, c3 chan int, c4 chan *doctor.Doctor) *WaitingRoom {
	once3.Do(func() {
		instance_wr = &WaitingRoom{
			MsgRequest:           make(chan *patient.Patient, 20),
			QueuePatients:        make([]*patient.Patient, 0),
			EmergencyRoomRequest: c1,
			EmergencyRoomReponse: c2,
			DoctorsRequest:       c3,
			DoctorsResponse:      c4,
			Alltime: map[string]float64{},
		}
	})
	return instance_wr
}

func (wr *WaitingRoom) handlerPatientWaitingRequest(p *patient.Patient) {
	wr.QueuePatients = append(wr.QueuePatients, p)
}

func (wr *WaitingRoom) Run() {
	log.Println("Waiting Room start working")
	for {
		select {
		case i := <-wr.MsgRequest:
			log.Println("Waiting room get a new patient " + strconv.FormatInt(int64(i.ID), 10) + " join request")
			wr.handlerPatientWaitingRequest(i)
		default:
			wr.work()
		}
	}
}

// Parcourir les patients, rechercher des patients disposant de ressources suffisantes pour le traitement, priorité de haut niveau
// pour un malade
// Demander un cours de gestion des urgences Demander des ressources pour les urgences
// Demander à la classe de gestion des médecins de demander des ressources médicales
// Les patients pouvant être traités sont exécutés par ordre de priorité
func (wr *WaitingRoom) work() {
	for i := 0; i < len(wr.QueuePatients); i++ {
		wr.DoctorsRequest <- wr.QueuePatients[i].Severity
		p1 := <-wr.DoctorsResponse
		wr.EmergencyRoomRequest <- wr.QueuePatients[i].Severity
		p2 := <-wr.EmergencyRoomReponse

		// Traite si les deux ressources sont disponibles en même temps
		if p1 != nil && p2 != nil {
			pp := wr.QueuePatients[i]
			p1.Lock()
			p1.Usable = false
			p1.Unlock()
			p2.Lock()
			p2.Status = 1
			p2.Unlock()
			pp.Status = patient.Being_treated_now
			wr.QueuePatients = append(wr.QueuePatients[:i], wr.QueuePatients[i+1:]...)
			go wr.treat(p2, p1, pp)
		}
	}
}

func (wr *WaitingRoom) treat(r *EmergencyRoom, d *doctor.Doctor, p *patient.Patient) {

	log.Println("patient " + strconv.FormatInt(int64(p.ID), 10) + " start to be treat by doctor " + strconv.FormatInt(int64(d.ID), 10) + " in room " + strconv.FormatInt(int64(r.ID), 10))
	time.Sleep(time.Duration(p.TimeForTreat) * time.Second)
	p.Status = patient.Finish
	log.Println("patient " + strconv.FormatInt(int64(p.ID), 10) + " finish treat ")

	wr.Alltime["patient " + strconv.FormatInt(int64(p.ID), 10) + "gravite " + strconv.FormatInt(int64(p.Severity), 10)] = time.Now().Sub(p.T).Seconds()

	r.Lock()
	r.Status = 0
	r.Unlock()
	d.Lock()
	d.Usable = true
	d.Unlock()
}
