package nurse

import (
	"errors"
	"log"
	"math/big"
	"strconv"
	"sync"
	"time"

	"crypto/rand"

	"gitlab.utc.fr/wanhongz/emergency-simulator/agent/patient"
)

type nurse struct {
	sync.Mutex
	ID       int                   // ID unique de l'infirmière
	Usable   bool                  // Est-ce actuellement gratuit
	p        *patient.Patient      // Patients actuellement sous jugement
	manager  *Nurse_manager        // La gestion
	msg_send chan *nurse           // Canal d'envoi de messages à la classe de gestion
	msg_recv chan *patient.Patient // Accepter les demandes de classe de gestion
}

// Constructeur
func NewNurse(id int, m *Nurse_manager) *nurse {
	return &nurse{
		ID:       id,
		Usable:   false,
		p:        nil,
		manager:  m,
		msg_send: m.msg_nurse,
		msg_recv: make(chan *patient.Patient, 20),
	}
}

func (n *nurse) GetChan() chan *patient.Patient {
	return n.msg_recv
}

// accepter de nouveaux patients
func (n *nurse) TreatNewPatient(p *patient.Patient) error {
	n.Lock()
	if n.Usable != true {
		n.Unlock()
		return errors.New("Patient is not usable")
	} else {
		n.p = p
		n.Unlock()
		return nil
	}
}

func (n *nurse) SetPatient(patient2 *patient.Patient) {
	n.p = patient2
}

// libérer
func (n *nurse) SetUsable(b bool) {
	n.Lock()
	n.Usable = b
	n.Unlock()
}

// définir le statut du patient
func (n *nurse) SetPatientStatus(gravite int, time int) {
	n.Lock()
	n.p.SetSeverity(gravite)
	n.p.SetTimeForTreat(time)
	n.Unlock()
}

// Informez le responsable que vous êtes libre et demandez une affectation
func (n *nurse) ticket() {
	n.msg_send <- n
}

// Diagnostiquer le patient
func (n *nurse) judge(patient2 *patient.Patient) {
	// Algorithme de jugement

	// temps de traitement de la simulation de sommeil
	// n.Lock()
	n.Lock()
	patient2.SetStatus(patient.Being_judged_by_nurse)
	timee, _ := rand.Int(rand.Reader, big.NewInt(5))
	tt := int(timee.Int64())
	tt += 3
	time.Sleep(time.Duration(tt) * time.Second)

	// sévérité
	gra, _ := rand.Int(rand.Reader, big.NewInt(4))

	// temp
	tim, _ := rand.Int(rand.Reader, big.NewInt(6))
	n.Unlock()
	if patient2.Severity == -1 {
		n.SetPatientStatus(int(gra.Int64()+1), 10+int(tim.Int64()))
	} else {
		n.SetPatientStatus(patient2.Severity, 5 + 2 * patient2.Severity + int(tim.Int64()))
	}

	patient2.Lock()
	patient2.Msg_nurse <- "ticket"
	patient2.Unlock()
	// n.Unlock()
}

func (nur *nurse) treat(n *patient.Patient) {
	nur.SetPatient(n)
	nur.SetUsable(false)
	// Appeler la fonction infirmière pour régler l'état du patient
	nur.judge(n)
	nur.ticket()
}

func (nur *nurse) Run() {
	log.Println("Nurse " + strconv.FormatInt(int64(nur.ID), 10) + " start")
	for {
		select {
		case n, ok := <-nur.msg_recv:
			if !ok {
				log.Println("Nurse " + strconv.FormatInt(int64(nur.ID), 10) + " stop")
				return
			}
			nur.treat(n)
		default:
			time.Sleep(1 * time.Second)
		}
	}
}
