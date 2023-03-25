package patient

import (
	"log"
	"strconv"
	"sync"
	"time"
)

type Patient struct {
	sync.Mutex
	ID                    int            // identifiant unique du patient
	Age                   int            // âge
	Gender                bool           // sexe (0 homme 1 femme)
	Symptom               string         // symptôme
	Severity              int            // échelle de symptômes 1-10
	Tolerance             int            // temps d'attente maximum (dépassant les congés/décès)
	TimeForTreat          int            // Temps de traitement estimé (minutes)
	Status                patient_status // État actuel et progrès du patient
	Msg_nurse             chan string    // canal pour recevoir les messages des infirmières
	Msg_request_nurse     chan *Patient  // demande le canal de l'infirmière
	Msg_request_reception chan *Patient  // canal pour demander l'enregistrement
	Msg_receive_reception chan string    // canal qui acceptent l'inscription
	Msg_request_waiting   chan *Patient  // Canal pour demander à rejoindre la file d'attente de la salle d'attente
	T                     time.Time      // temps de creation
}

// Constructeur
func NewPatient(id int, age int, gender bool, symptom string, severity int, tolerance int, timeForTreate int, c chan *Patient, d chan *Patient, c_w chan *Patient) *Patient {
	return &Patient{
		ID:                    id,
		Age:                   age,
		Gender:                gender,
		Symptom:               symptom,
		Severity:              severity,
		Tolerance:             tolerance,
		TimeForTreat:          timeForTreate,
		Status:                Waiting_for_nurse,
		Msg_nurse:             make(chan string, 20),
		Msg_request_nurse:     c,
		Msg_request_reception: d,
		Msg_receive_reception: make(chan string, 20),
		Msg_request_waiting:   c_w,
	}
}

func (p *Patient) SetSeverity(s int) {
	p.Lock()
	p.Severity = s
	p.Unlock()
}

func (p *Patient) SetTimeForTreat(s int) {
	p.Lock()
	p.TimeForTreat = s
	p.Unlock()
}

func (p *Patient) SetStatus(s patient_status) {
	p.Lock()
	p.Status = s
	p.Unlock()
}

func (p *Patient) RequestCheckingStatus() {
	p.Lock()
	p.Msg_request_nurse <- p
	p.Unlock()
}

func (p *Patient) Run() {
	p.RequestCheckingStatus()
	for {
		select {
		case n := <-p.Msg_nurse:
			if n == "ticket" {
				p.SetStatus(Waiting_for_register)
				log.Println("Patient " + strconv.FormatInt(int64(p.ID), 10) + " get a status: gravity " + strconv.FormatInt(int64(p.Severity), 10) + ", need time " + strconv.FormatInt(int64(p.TimeForTreat), 10) + ", and go to get a reception")
				p.Lock()
				p.Msg_request_reception <- p
				p.Unlock()
			}
		case m := <-p.Msg_receive_reception:
			if m == "ticket" {
				log.Println("Patient " + strconv.FormatInt(int64(p.ID), 10) + " get reception")
				p.SetStatus(Waiting_for_treat)
				p.Lock()
				p.Msg_request_waiting <- p
				p.Unlock()
			}

		default:
			time.Sleep(1 * time.Second)
		}
	}
}
